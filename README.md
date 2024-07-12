"# Go-Gin-_REST" 

git add .
git commit -m "go rest commit"
git push -u origin main

go mod init testing/new

go get
go mod tidy

go run main.go

docker pull mysql:latest
docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=mysql#pass -p 3306:3306 -d mysql:latest
docker exec -it test-mysql bash
mysql -u root -p

CREATE DATABASE test;
use test;

CREATE TABLE posts (
  id serial primary key,
  username varchar(255) not null,
  title varchar(100) not null,
  content text not null,
  created_at timestamptz not null default clock_timestamp(),
  updated_at timestamptz
);

CREATE TABLE quote
( id INT(11) NOT NULL AUTO_INCREMENT , 
 quote VARCHAR(255) NOT NULL , 
 author VARCHAR(255) NOT NULL , 
 created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP , 
 updated_at DATETIME on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , 
 PRIMARY KEY (id), 
 INDEX idx_author (author), UNIQUE idx_quote_uniqie (quote)
) 
 ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;

INSERT INTO quote (id, quote, author) VALUES
(1, 'There are only two kinds of languages: the ones people complain about and the ones nobody uses.', 'Bjarne Stroustrup'),
(3, 'Any fool can write code that a computer can understand. Good programmers write code that humans can understand.', 'Martin Fowler'),
(4, 'First, solve the problem. Then, write the code.', 'John Johnson'),
(5, 'Java is to JavaScript what car is to Carpet.', 'Chris Heilmann'),
(6, 'Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live.', 'John Woods'),
(7, "I'm not a great programmer; I'm just a good programmer with great habits.", 'Kent Beck'),
(8, 'Truth can only be found in one place: the code.', 'Robert C. Martin'),
(9, "If you have to spend effort looking at a fragment of code and figuring out what it's doing, then you should extract it into a function and name the function after the 'what'.", 'Martin Fowler'),
(10, 'The real problem is that programmers have spent far too much time worrying about efficiency in the wrong places and at the wrong times; premature optimization is the root of all evil (or at least most of it) in programming.', 'Donald Knuth'),
(11, 'SQL, Lisp, and Haskell are the only programming languages that I’ve seen where one spends more time thinking than typing.', 'Philip Greenspun'),
(12, 'Deleted code is debugged code.', 'Jeff Sickel'),
(13, 'There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies and the other way is to make it so complicated that there are no obvious deficiencies.', 'C.A.R. Hoare'),
(14, 'Simplicity is prerequisite for reliability.', 'Edsger W. Dijkstra'),
(15, 'There are only two hard things in Computer Science: cache invalidation and naming things.', 'Phil Karlton'),
(16, 'Measuring programming progress by lines of code is like measuring aircraft building progress by weight.', 'Bill Gates');

SELECT * FROM quote;


CREATE TABLE IF NOT EXISTS users-test (
id int(11) NOT NULL AUTO_INCREMENT ,
name varchar(200) NOT NULL,
description varchar(200) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
ALTER TABLE users ADD PRIMARY KEY (id);
ALTER TABLE users MODIFY id int(11) NOT NULL AUTO_INCREMENT;

INSERT INTO users (id, name, email, created_at) VALUES
  (1, 'Test', 'test@g.co', '2019-02-28 13:20:20'),
  (2, 'john', 'john@g.co', '2019-02-28 13:20:20'),
  (3, 'tutsmake', 'tuts@g.co', '2019-02-28 13:20:20'),
  (4, 'tut', 'tut@g.co', '2019-02-28 13:20:20'),
  (5, 'mhd', 'mhd@g.co', '2019-02-28 13:20:20');

SELECT * FROM users;


CREATE TABLE IF NOT EXISTS user_test (
id int(11) NOT NULL AUTO_INCREMENT ,
name varchar(200) NOT NULL,
description varchar(200) NOT NULL,
PRIMARY KEY (id)
) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;

INSERT INTO user_test (id, name, description) VALUES
  (1, 'Test no.1', 'king'),
  (2, 'Test no.2', 'queen'),
  (3, 'Test no.3', 'jack');

SELECT * FROM user_test;

CREATE DATABASE movies;
use movies;
CREATE TABLE movies (  
  id int(5) DEFAULT NULL, 
  title varchar(255) DEFAULT NULL, 
  director varchar(255) DEFAULT NULL ) 
  ENGINE=InnoDB DEFAULT CHARSET= utf8;

DESCRIBE movies;

INSERT INTO movies VALUES (
  '01','jurrasic Park','steven Spielberg');

SELECT * FROM movies

curl http://localhost:8080/users \
    --include \
    --header "Content-Type: application/json" \
    --request "GET"

curl http://localhost:8080/users \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "5","name": "test no.5","description": "nothing"}'