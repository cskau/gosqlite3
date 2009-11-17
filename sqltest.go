package main

import "fmt"
import "sqlite3"

func main() {
	sqlite3.Initialize();
	defer sqlite3.Shutdown();
	fmt.Printf("Sqlite3 Version: %s\n", sqlite3.LibVersion());

	dbh := new(sqlite3.Handle);
	dbh.Open("test.db");
	defer dbh.Close();

	st,err := dbh.Prepare("CREATE TABLE foo (i INTEGER, s VARCHAR(20));");
	if (err != "") {
		println(err);
	} else {
		defer st.Finalize();
		st.Step();
	}
	
	st,err = dbh.Prepare("INSERT INTO foo values (2, 'this is a test')");
	if (err != "") {
		println(err);
	} else {
		defer st.Finalize();
		st.Step();
	}

	st,err = dbh.Prepare("INSERT INTO foo values (3, 'holy moly')");
	if (err != "") {
		println(err);
	} else {
		defer st.Finalize();
		st.Step();
	}
	
	fmt.Printf("%d changes\n", dbh.TotalChanges());
	
	fmt.Printf("last insert id: %d\n", dbh.LastInsertRowID())
	
}
