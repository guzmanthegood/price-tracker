CREATE TABLE price (
    id              integer         NOT NULL,
    amount      	decimal(5,2)    NOT NULL,
	flight_number	varchar(255)			,
    origin          char(3)		    NOT NULL,
    destination     char(3)    		NOT NULL,
    departure       date            		,
    comeback        date            		,
	oneway			bool			NOT NULL
);