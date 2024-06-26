-- チャット
CREATE TABLE `channel` (
    `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT "チャンネルID",
    `channel_key` VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "チャンネルKEY",
    `room_key`    VARCHAR(20)  NOT NULL                COMMENT "ルームKEY",
    `name`        VARCHAR(50)  NOT NULL                COMMENT "チャンネル名",
    `description` VARCHAR(191) NOT NULL                COMMENT "説明",
    `type`        VARCHAR(191) NOT NULL                COMMENT "種別",
    `deleted_at`  TIMESTAMP    DEFAULT NULL            COMMENT "削除日時",
    `created_at`  TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`  TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
