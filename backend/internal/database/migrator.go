//very important file 
package database

import(
	"context",
	"github.com/jackc/pgx/v5/pgxpool",
	"embed",
	"fmt",
	"io/fs",
	"sort",
	"strings"
)
//go:embed migrations/*.sql
//at build time take all the sql files inside the migrations folders and embed them into the binary
var migrationFiles embed.FS

func RunMigrations(ctx context.Context, db *pgxpool.Pool) error {
	//we have to create table to track that which migrations already been executed
	//to prevent same migrations will not run again
	if _, err := db.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
version VARCHAR PRIMARY KEY,
applied_at TIMESTAMP NOT NULL DEFAULT NOW()
)`); err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	entries, err := fs.ReadDir(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %w", err)
	}
	var filenames []string
	//we are only going to see the sql files
	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".sql") {
			filenames = append(filenames, entry.Name())
		}
	}
	sort.Strings(filenames) //sort the files alphabetically to ensure they run in the correct order
	for_,filename:=range filenames{
		var alreadyApplied bool
  //scan works like rows
		if err:=db.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`, filename).Scan(&alreadyApplied); err!=nil{
			return fmt.Errorf("failed to check migration version: %w", err)
		}

		if alreadyApplied {	continue}
		sqlBytes, err := migrationFiles.ReadFile("migrations/" + filename)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", filename, err)
		}
		transactions,err:=db.Begin(ctx)
		if err!=nil{
			return fmt.Errorf("failed to begin transaction for migration %s: %w", filename, err)
		}
		if _,err:=transactions.Exec(ctx, string(sqlBytes)); err!=nil{
			_=transactions.Rollback(ctx)
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}
		if _,err:=transactions.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, filename); err!=nil{
			_=transactions.Rollback(ctx)
			return fmt.Errorf("failed to record migration version %s: %w", filename, err)
		}
		if err:=transactions.Commit(ctx); err!=nil{
			return fmt.Errorf("failed to commit transaction for migration %s: %w", filename, err)
		}
	}
}
