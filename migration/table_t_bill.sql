CREATE TABLE t_bill (
    id BIGINT NOT NULL AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    total_price INT NOT NULL DEFAULT 0,
    bill_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (id),
    PRIMARY KEY (id)
);

ALTER TABLE t_bill
ADD CONSTRAINT fk_bill_order FOREIGN KEY (order_id) REFERENCES t_order(id);