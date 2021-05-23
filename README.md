# order-service
This project is simple checkout system

## Prerequisite
Please provide postgres db connection, and run query bellow:

Create Table
```
CREATE TABLE public.product (
	id serial NOT NULL,
	sku varchar(255) NOT NULL,
	product_name varchar(255) NULL,
	price float8 NULL,
	quantity int4 NULL,
	promo_code varchar(10) NULL,
	CONSTRAINT product_pkey PRIMARY KEY (id),
	CONSTRAINT product_sku_key UNIQUE (sku)
);

CREATE TABLE public.promo (
	id serial NOT NULL,
	code varchar(255) NOT NULL,
	discount float8 NULL,
	free_items int4 NULL,
	minimum_buy int4 NULL,
	CONSTRAINT promo_code_key UNIQUE (code),
	CONSTRAINT promo_pkey PRIMARY KEY (id)
);

ALTER TABLE public.promo ADD CONSTRAINT fk_product FOREIGN KEY (free_items) REFERENCES product(id);

CREATE TABLE public.users (
	id int4 NOT NULL,
	username varchar(255) NOT NULL,
	address varchar(255) NULL,
	CONSTRAINT user_username_key UNIQUE (username),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE public.cart (
	id serial NOT NULL,
	user_id int4 NOT NULL,
	product_id int4 NOT NULL,
	total float8 NOT NULL,
	quantity int4 NOT NULL,
	description varchar NULL,
	free_items int4 NULL,
	total_free_items int4 NULL,
	status varchar NULL,
	CONSTRAINT cart_pkey PRIMARY KEY (id)
);

ALTER TABLE public.cart ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product(id);
ALTER TABLE public.cart ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id);
```

Insert Sample Data
```
INSERT INTO public.product (sku,product_name,price,quantity,promo_code) VALUES
	 ('43N23P','Macbook Pro',5399.99,5,'PROMO1'),
	 ('A304SD','Alexa Speaker',109.5,10,'PROMO3'),
	 ('234234','Raspberry Pi B',30.0,2,NULL),
	 ('120P90','Google Home',49.99,10,'PROMO2');


INSERT INTO public.promo (code,discount,free_items,minimum_buy) VALUES
	 ('PROMO2',33.0,NULL,3),
	 ('PROMO3',10.0,NULL,3),
	 ('PROMO1',NULL,4,1);


INSERT INTO public.users (id,username,address) VALUES
	 (2,'Fathan','klaten');
```

## How To Run
1. change db configuration to your db connection in /database/init.go:
```
const (
	host     = "{your_db_host}"
	port     = {your_db_port}
	user     = "{your_db_user}"
	password = "{your_db_password}"
	dbname   = "{your_db_name}"
)
```
2. Simply run golang app using this command : `go run app.go`, and open `localhost:9011`, your ready to go!

## Sample Query
### Get All Product
```
{
  products(){
    sku
    name
    price
    promoCode
    quantity
  }
}
```
### Get Product By ID
```
{
  product(id: "1"){
    sku
    name
    price
    promoCode
    quantity
  }
}
```
### Get All Cart
```
{
  carts(){
    userId
    productId
    total
    freeItems
    totalFreeItems
    description
    quantity
    status
  }
}
```
total : quantity * price * discount

freeItems : bonus items promo

totalFreeItems : how many user get free items

### Add to cart
```
{
  addToCart(productID: 1, userID: 2, quantity: 3){
    userId
    productId
    total
    freeItems
    totalFreeItems
    description
    quantity
    status
  }
}
```

### Checkout
```
{
  checkout(userID: 2, pay: 100){
    totalOrder
    message
    isSuccess
  }
}
```