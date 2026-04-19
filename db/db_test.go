package db_test

import (
	"database/sql"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	_ "modernc.org/sqlite"
)

// openTestDB creates a temp SQLite DB and applies all *.up.sql migrations in order.
// `go test ./db/...` sets CWD to the db/ directory, so migrations/ is a sibling.
func openTestDB(t *testing.T) *sql.DB {
	t.Helper()

	tmpFile := filepath.Join(t.TempDir(), "test.db")
	db, err := sql.Open("sqlite", tmpFile)
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	entries, err := os.ReadDir("migrations")
	if err != nil {
		t.Fatalf("read migrations dir: %v", err)
	}

	var upFiles []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".up.sql") {
			upFiles = append(upFiles, e.Name())
		}
	}
	sort.Strings(upFiles)

	for _, name := range upFiles {
		sql, err := os.ReadFile(filepath.Join("migrations", name))
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		if _, err := db.Exec(string(sql)); err != nil {
			t.Fatalf("exec %s: %v", name, err)
		}
	}

	return db
}

func TestAgentsTableExists(t *testing.T) {
	db := openTestDB(t)

	var name string
	err := db.QueryRow(
		"SELECT name FROM sqlite_master WHERE type='table' AND name='agents'",
	).Scan(&name)
	if err != nil {
		t.Fatalf("agents table not found: %v", err)
	}
	if name != "agents" {
		t.Fatalf("expected table name 'agents', got %q", name)
	}
}

func TestAgentsTableColumns(t *testing.T) {
	db := openTestDB(t)

	rows, err := db.Query("PRAGMA table_info(agents)")
	if err != nil {
		t.Fatalf("PRAGMA table_info: %v", err)
	}
	defer rows.Close()

	wantCols := map[string]bool{
		"id":            false,
		"name":          false,
		"email":         false,
		"password_hash": false,
		"created_at":    false,
	}

	for rows.Next() {
		var cid int
		var colName, colType string
		var notNull, pk int
		var dfltValue sql.NullString
		if err := rows.Scan(&cid, &colName, &colType, &notNull, &dfltValue, &pk); err != nil {
			t.Fatalf("scan: %v", err)
		}
		if _, ok := wantCols[colName]; ok {
			wantCols[colName] = true
		}
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows: %v", err)
	}

	for col, found := range wantCols {
		if !found {
			t.Errorf("expected column %q not found in agents table", col)
		}
	}
}

func TestAgentsTableInsertAndQuery(t *testing.T) {
	db := openTestDB(t)

	_, err := db.Exec(
		"INSERT INTO agents (name, email, password_hash) VALUES (?, ?, ?)",
		"TestBot", "test@agentclinic.io", "hashedpassword",
	)
	if err != nil {
		t.Fatalf("insert: %v", err)
	}

	var email string
	err = db.QueryRow("SELECT email FROM agents WHERE name = ?", "TestBot").Scan(&email)
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	if email != "test@agentclinic.io" {
		t.Errorf("expected email %q, got %q", "test@agentclinic.io", email)
	}
}
