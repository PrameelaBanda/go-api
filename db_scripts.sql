CREATE TABLE deliveries (
   ID           int primary key not null,
   ORDER_ITEM   int                     not null,
   DELIVERED_QUANTITY        int        not null
);

CREATE TABLE order_items (
    ID           int primary key not null,
    ORDER_ID   int                     not null,
    PRICE_PER_UNIT        float8        not null,
    QUANTITY   int         not null,
    PRODUCT    varchar(64) not null
);

CREATE TABLE orders (
    ID int primary key not null,
    CREATED_AT timestamp not null,
    ORDER_NAME varchar(64) not null,
    CUSTOMER_ID varchar(64) not null
);

CREATE TABLE customer_companies (
    COMPANY_ID int primary key not null,
    COMPANY_NAME varchar(128) not null
);

CREATE TABLE customers (
    USER_ID varchar(64) primary key not null,
    LOGIN varchar(64) not null,
    PASSWORD varchar(64) not null,
    NAME varchar(128) not null,
    COMPANY_ID int not null,
    CREDIT_CARDS varchar ARRAY null
);