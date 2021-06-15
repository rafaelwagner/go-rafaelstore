GRANT ALL PRIVILEGES ON DATABASE rafaelstore TO postgres;

\c rafaelstore

CREATE TABLE products
(
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    description character varying(100),
    price numeric NOT NULL,
    stock integer NOT NULL
);

insert into products(id, name, description, price, stock) values(1, 'Test 1', 'Nice', 5, 99);
insert into products(id, name, description, price, stock) values(2, 'Test 2', 'Nice', 20, 50);