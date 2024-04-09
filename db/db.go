package db

import (
	// "database/sql"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	DB, err = sqlx.Open("mysql", "root:Sutharj@571@tcp(127.0.0.1:3306)/test1?parseTime=true")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(255) NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table: " + err.Error())
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTO_INCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
    )
    `

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table: " + err.Error())
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Could not create registrations table: " + err.Error())
	}

	createGssWishlistTable := `
	CREATE TABLE IF NOT EXISTS wishlistGssTBL (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		user_Id INT NOT NULL,
		space_Id INT UNSIGNED NOT NULL,
		start_Date DATE NOT NULL,
		end_Date DATE NOT NULL,
		is_add BIT(1),
		created_On timestamp not null default current_timestamp,
		updated_On timestamp not null default current_timestamp on update current_timestamp,
		PRIMARY KEY (id),
		CONSTRAINT fk_user FOREIGN KEY (user_Id) REFERENCES users(id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPRESSED;
	`
	_, err = DB.Exec(createGssWishlistTable)

	if err != nil {
		panic("Could not create GSSWishlist table: " + err.Error())
	}

	createGssTable := `
	create table if not exists gssTBL (
		gssID int unsigned auto_increment,
		uuid  varchar(40) not null unique,
		gssTypeID tinyint unsigned not null,
		locID int unsigned not null default 0,
		userID int unsigned not null,
		gssName varchar(100) not null,
		gssDesc varchar(300) not null default '',
		isShareable tinyint unsigned not null default 0,  -- 1: independent, 2: shareable
		isPOBF boolean not null default false,
		cnt int unsigned not null default 1,
		-- avlFrom date not null default(current_date), -- available from date
		avlFrom       date, -- available from date (curdate())
		maxBookDur tinyint unsigned not null default 0,  -- 1 to 12 months
		gssAccessID tinyint unsigned not null default 0, -- 1: 24 hrs intimation, 2: operational hours,
		bookingTypeID tinyint unsigned not null default 0, -- 1: instant: host approval not req, 2: host approval req
		len int unsigned not null default 0, -- in feet    -- length is a function name in mysql.
		breadth int unsigned not null default 0, -- in feet
		height int unsigned not null default 0, -- in feet
		area int unsigned not null default 0, -- in sq. feet
		coverImageURL varchar(512) not null default '', -- fqdn of cover-image.
		-- gssUsageID tinyint unsigned not null,
	
		-- status
		isPublished boolean not null default false,
		isDisabled  boolean not null default false,
		isDeleted boolean not null default false,
	
		-- audit
		createdOn timestamp not null default current_timestamp,
		updatedOn timestamp not null default current_timestamp on update current_timestamp,
	
		-- constraints
		-- primary key(s)
		primary key(gssID)

	)	engine = innodb default charset = utf8mb4 collate = utf8mb4_unicode_ci row_format = compressed;
	`

	_, err = DB.Exec(createGssTable)

	if err != nil {
		panic("Could not create GSSWishlist table: " + err.Error())
	}

}
