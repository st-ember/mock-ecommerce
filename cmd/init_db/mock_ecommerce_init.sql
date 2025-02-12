CREATE TABLE "purchase"(
    "id" UUID NOT NULL,
    "product_id" UUID NOT NULL,
    "customer_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "customer"(
    "id" UUID NOT NULL,
    "username" VARCHAR(50) NOT NULL,
    "region_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "product_category"(
    "id" UUID NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "slug" VARCHAR(50) NOT NULL
);

CREATE TABLE "product"(
    "id" UUID NOT NULL,
    "category_id" UUID NOT NULL,
    "product_name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "merchant_id" UUID NOT NULL,
    "price" NUMERIC(12, 2) NOT NULL
);

CREATE TABLE "region"(
    "id" UUID NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "code" CHAR(3) NOT NULL
);

CREATE TABLE "discount_event"(
    "id" UUID NOT NULL,
    "sale_category_id" UUID NOT NULL,
    "merchant_id" UUID NOT NULL,
    "start_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "end_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "product_category_id" UUID NOT NULL
);

CREATE TABLE "merchant"(
    "id" UUID NOT NULL,
    "username" VARCHAR(50) NOT NULL,
    "region_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "discount_category"(
    "id" UUID NOT NULL,
    "discount_percentage" INT NOT NULL,
    "code" CHAR(5) NOT NULL
);

// pk
ALTER TABLE "purchase" ADD PRIMARY KEY("id");
ALTER TABLE "customer" ADD PRIMARY KEY("id");
ALTER TABLE "product_category" ADD PRIMARY KEY("id");
ALTER TABLE "product" ADD PRIMARY KEY("id");
ALTER TABLE "region" ADD PRIMARY KEY("id");
ALTER TABLE "discount_event" ADD PRIMARY KEY("id");
ALTER TABLE "merchant" ADD PRIMARY KEY("id");
ALTER TABLE "discount_category" ADD PRIMARY KEY("id");

// fk
ALTER TABLE "purchase" ADD CONSTRAINT "purchase_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "customer"("id");
ALTER TABLE "customer" ADD CONSTRAINT "customer_region_id_foreign" FOREIGN KEY("region_id") REFERENCES "region"("id");
ALTER TABLE "discount_event" ADD CONSTRAINT "discount_event_sale_category_id_foreign" FOREIGN KEY("sale_category_id") REFERENCES "discount_category"("id");
ALTER TABLE "purchase" ADD CONSTRAINT "purchase_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "product"("id");
ALTER TABLE "merchant" ADD CONSTRAINT "merchant_region_id_foreign" FOREIGN KEY("region_id") REFERENCES "region"("id");
ALTER TABLE "discount_event" ADD CONSTRAINT "discount_event_merchant_id_foreign" FOREIGN KEY("merchant_id") REFERENCES "merchant"("id");
ALTER TABLE "discount_event" ADD CONSTRAINT "discount_event_product_category_id_foreign" FOREIGN KEY("product_category_id") REFERENCES "product_category"("id");
ALTER TABLE "product" ADD CONSTRAINT "product_category_id_foreign" FOREIGN KEY("category_id") REFERENCES "product_category"("id");
ALTER TABLE "product" ADD CONSTRAINT "product_merchant_id_foreign" FOREIGN KEY("merchant_id") REFERENCES "merchant"("id");