package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"stock-bit/config"
)

type Repositories struct {
	db    *sql.DB
	Movie MovieRepository
}

func NewRepositories(configure config.Configuration) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s",
		configure.DBHost, configure.DBPort, configure.DBUser, configure.DBName, configure.DBPassword)
	db, err := sql.Open(configure.DBDriver, DBURL)
	err = db.Ping()
	if err != nil {
		fmt.Printf("Cannot connect to %s database url %s", configure.DBDriver, DBURL)
		log.Println("\nThis is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database with url %s\n", configure.DBDriver, DBURL)
	}
	return &Repositories{
		db:    db,
		Movie: NewMovieRepository(db),
	}, nil
}

// Close closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// Seeder This migrate all tables
func (s *Repositories) Seeder() error {
	var err error
	queryMovies := fmt.Sprintf("create table if not exists movies (id serial primary key, title varchar(255), year varchar(255), imdb_id varchar(255), type varchar(255), poster varchar(255);")
	_, err = s.db.Exec(queryMovies)
	queryMoviesDetail := fmt.Sprintf("create table if not exists movies_detail (id serial primary key, title varchar(55) unique, year varchar(55), " +
		"rated varchar(55), released varchar(55), runtime varchar(55), genre varchar(55), director varchar(55), writer varchar(55), actors text, plot text, " +
		"language varchar(55), country varchar(255), awards text, poster varchar(255));")
	_, err = s.db.Exec(queryMoviesDetail)
	return err
}

func (s *Repositories) AddForeignKey() error {
	var err error
	return err
}
