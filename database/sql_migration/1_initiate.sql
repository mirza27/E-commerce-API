-- +migrate Up
-- +migrate StatementBegin

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE customer(
    id SERIAL PRIMARY KEY NOT NULL,
    uname VARCHAR(255),
    password VARCHAR(255),
    cash BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE category(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(256),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE item(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(256),
    description VARCHAR(256),
    stock INT NOT NULL,
    price BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category(id) 
);


CREATE TABLE order_history(
    id SERIAL PRIMARY KEY NOT NULL,
    item_id BIGINT NOT NULL,
    customer_id BIGINT NOT NULL,
    number_of_item BIGINT NOT NULL,
    bill BIGINT NOT NULL,
    Purchased_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES item(id),
    CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customer(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON order_history
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON item
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON customer
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


-- +migrate StatementEnd
