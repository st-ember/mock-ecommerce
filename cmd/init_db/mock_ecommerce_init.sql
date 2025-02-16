DO $$ 
DECLARE 
  r RECORD;
BEGIN
  -- Drop all tables
  FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
    EXECUTE 'DROP TABLE IF EXISTS public.' || r.tablename || ' CASCADE';
  END LOOP;


  -- Drop all indexes
  FOR r IN (SELECT indexname FROM pg_indexes WHERE schemaname = 'public') LOOP
    EXECUTE 'DROP INDEX IF EXISTS public.' || r.indexname || ' CASCADE';
  END LOOP;
END $$;

CREATE TABLE "customer"(
    "id" UUID NOT NULL,
    "username" VARCHAR(50) NOT NULL,
    "country" UUID NOT NULL,
    "joined_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "merchant"(
    "id" UUID NOT NULL,
    "username" VARCHAR(50) NOT NULL,
    "country" UUID NOT NULL,
    "description" TEXT NOT NULL,
    "status" INT NOT NULL,
    "joined_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "purchase"(
    "id" UUID NOT NULL,
    "product" UUID NOT NULL,
    "customer" UUID NOT NULL,
    "purchased_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "discount_event"(
    "id" UUID NOT NULL,
    "discount_detail" UUID NOT NULL,
    "product_category" UUID NOT NULL,
    "merchant" UUID NOT NULL,
    "start_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "end_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "discount_detail"(
    "id" UUID NOT NULL,
    "is_percentage" BOOLEAN NOT NULL,
    "percentage_value" NUMERIC(2, 2) NOT NULL,
    "is_get_n_free" BOOLEAN NOT NULL,
    "buy_ammount" INTEGER NOT NULL,
    "free_ammount" INTEGER NOT NULL
);

CREATE TABLE "product"(
    "id" UUID NOT NULL,
    "category" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "merchant" UUID NOT NULL,
    "price" NUMERIC(12, 2) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "product_category"(
    "id" UUID NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "slug" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "country"(
    "id" UUID NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "code" VARCHAR(50) NOT NULL
);

-- setting primary keys
ALTER TABLE "customer" ADD PRIMARY KEY("id");
ALTER TABLE "merchant" ADD PRIMARY KEY("id");
ALTER TABLE "purchase" ADD PRIMARY KEY("id");
ALTER TABLE "country" ADD PRIMARY KEY("id");
ALTER TABLE "product" ADD PRIMARY KEY("id");
ALTER TABLE "product_category" ADD PRIMARY KEY("id");
ALTER TABLE "discount_event" ADD PRIMARY KEY("id");
ALTER TABLE "discount_detail" ADD PRIMARY KEY("id");

-- setting foreign keys
ALTER TABLE "purchase" ADD CONSTRAINT "purchase_customer_foreign" FOREIGN KEY("customer") REFERENCES "customer"("id");
ALTER TABLE "customer" ADD CONSTRAINT "customer_country_foreign" FOREIGN KEY("country") REFERENCES "country"("id");
ALTER TABLE "discount_event" ADD CONSTRAINT "discount_event_discount_detail_foreign" FOREIGN KEY("discount_detail") REFERENCES "discount_detail"("id");
ALTER TABLE "purchase" ADD CONSTRAINT "purchase_product_foreign" FOREIGN KEY("product") REFERENCES "product"("id");
ALTER TABLE "merchant" ADD CONSTRAINT "merchant_country_foreign" FOREIGN KEY("country") REFERENCES "country"("id");
ALTER TABLE "discount_event" ADD CONSTRAINT "discount_event_merchant_foreign" FOREIGN KEY("merchant") REFERENCES "merchant"("id");
ALTER TABLE "discount_event" ADD CONSTRAINT "discount_event_product_category_foreign" FOREIGN KEY("product_category") REFERENCES "product_category"("id");
ALTER TABLE "product" ADD CONSTRAINT "product_category_foreign" FOREIGN KEY("category") REFERENCES "product_category"("id");
ALTER TABLE "product" ADD CONSTRAINT "product_merchant_foreign" FOREIGN KEY("merchant") REFERENCES "merchant"("id");