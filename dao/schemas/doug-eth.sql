-- Blocks
create table blocks
(
    block_id    serial
        constraint blocks_pk
            primary key,
    block_num   int not null,
    block_hash  text,
    block_time  timestamp,
    parent_hash text
);

-- Transactions
create table transactions
(
    transaction_id   integer not null
        constraint transactions_pk
            primary key,
    block_id         integer
        constraint transactions_blocks_block_id_fk
            references blocks,
    transaction_hash text
);
