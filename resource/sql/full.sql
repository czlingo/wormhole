DROP TABLE IF EXISTS pipeline;
CREATE TABLE pipeline (
    `id`            VARCHAR(64)         PRIMARY KEY,
    `name`          VARCHAR(64)         NOT NULL,
    `project`       VARCHAR(64)         NOT NULL,
    `content`       VARCHAR(512)        NOT NULL,
    `create_time`   BIGINT              NOT NULL,
    `update_time`   BIGINT              NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;