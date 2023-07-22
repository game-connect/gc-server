CREATE TABLE `user` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "ユーザーID",
    `user_key` VARCHAR(12) NOT NULL UNIQUE COMMENT "ユーザーKey",
    `username` VARCHAR(50) NOT NULL COMMENT "ユーザー名",
    `email` VARCHAR(255) NOT NULL COMMENT "メールアドレス",
    `password` VARCHAR(255) NOT NULL COMMENT "パスワード",
    `token` VARCHAR(255) NOT NULL COMMENT "アクセストークン",
    `status` VARCHAR(255) NOT NULL COMMENT "状態",
    `created_at` TIMESTAMP NOT NULL COMMENT "作成日時",
    `updated_at` TIMESTAMP NOT NULL COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


