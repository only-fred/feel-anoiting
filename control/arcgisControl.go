package control

import (
	"fmt"
	"log"

	conn "../connection/db"
)

func Create(provincestate string, countryregion string, lastupdate string, confirmed int, recovered int, deaths int, active int) {
	db := conn.Connection()

	result, err := db.Exec("INSERT INTO tbl_attributes (province_state, country_region, last_update, confirmed, recovered, deaths, active) VALUES ($1, $2, $3, $4, $5, $6, $7)", provincestate, countryregion, lastupdate, confirmed, recovered, deaths, active)
	if err != nil {
		log.Fatal("Failed to INSERT ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Failed to read rows affected ", err)
	}

	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d", rows)
	}

	db.Close()
}

func Read() string {
	db := conn.Connection()

	rows, err := db.Query("SELECT*FROM tbl_attributes")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id            int
			provincestate string
			countryregion string
			lastupdate    string
			confirmed     int
			recovered     int
			deaths        int
			active        int
		)
		if err := rows.Scan(&id, &provincestate, &countryregion, &lastupdate, &confirmed, &recovered, &deaths, &active); err != nil {
			log.Fatal(err)
		}

		return lastupdate
	}

	db.Close()

	return ""
}

func ReadAll() {
	db := conn.Connection()

	rows, err := db.Query("SELECT*FROM tbl_attributes")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n----+----------------+----------------+-------------+-----------+-----------+--------+--------")

	for rows.Next() {
		var (
			id            int
			provincestate string
			countryregion string
			lastupdate    string
			confirmed     int
			recovered     int
			deaths        int
			active        int
		)
		if err := rows.Scan(&id, &provincestate, &countryregion, &lastupdate, &confirmed, &recovered, &deaths, &active); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("| ID: %d \n| State: %s \n| Country: %s \n| Last Update: %s \n| Confirmed: %d \n| Recovered: %d \n| Deaths: %d \n| Active: %d ", id, provincestate, countryregion, lastupdate, confirmed, recovered, deaths, active)
		fmt.Println("\n----+----------------+----------------+-------------+-----------+-----------+--------+--------")
	}

	db.Close()
}

func Update(provincestate string, countryregion string, lastupdate string, confirmed int, recovered int, deaths int, active int, id int) {
	db := conn.Connection()

	result, err := db.Exec("UPDATE tbl_attributes SET province_state=$1, country_region=$2, last_update=$3, confirmed=$4, recovered=$5, deaths=$6, active=$7 WHERE id=$8",
		provincestate, countryregion, lastupdate, confirmed, recovered, deaths, active, id)
	if err != nil {
		log.Fatal("Failed to UPDATE ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Failed to read rows affected ", err)
	}

	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d", rows)
	}

	db.Close()
}
