-- CreateTable
CREATE TABLE `user` (
                        `id` VARCHAR(191) NOT NULL,
                        `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
                        `email` VARCHAR(191) NOT NULL,
                        `name` VARCHAR(191) NOT NULL,
                        `password` VARCHAR(191) NOT NULL,
                        `roleId` VARCHAR(191) NOT NULL,

                        UNIQUE INDEX `user_email_key`(`email`),
                        PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `role` (
                        `id` VARCHAR(191) NOT NULL,
                        `name` VARCHAR(191) NOT NULL,

                        UNIQUE INDEX `role_name_key`(`name`),
                        PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `user` ADD CONSTRAINT `user_roleId_fkey` FOREIGN KEY (`roleId`) REFERENCES `role`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- RoleSeeder
INSERT INTO role VALUES ('49b1d442-f912-4ef7-aa88-e889be7b6bda', 'user');
INSERT INTO role VALUES ('dec578b7-9785-4acd-9ddc-9a2d2ddabb26', 'admin');