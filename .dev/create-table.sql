create schema if not exists products;

create table if not exists products.stock_positions (
    id varchar(100) not null,
    item_id varchar(100) not null, 
    facility_id number(4), 
    on_hand_qty number(20),
    unavl_qty number(20),
    position_date timestamp(6),
    transaction_id varchar(100),
    primary key (id)
);

create user if not exists stock_pos_user_service identified by '123';
grant all on products.stock_positions to stock_pos_user_service;