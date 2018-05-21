CREATE TABLE price (
    id              serial          PRIMARY KEY,
    amount      	decimal(5,2)    NOT NULL,
	flight_number	varchar(20)			,
    origin          char(3)		    NOT NULL,
    destination     char(3)    		NOT NULL,
    departure       date            		,
    comeback        date            		,
	oneway			bool			NOT NULL,
    created_at      date            NOT NULL
);