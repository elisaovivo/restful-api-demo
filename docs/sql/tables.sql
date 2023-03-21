CREATE TABLE `resource` (
                            `id` char(64) CHARACTER SET latin1 NOT NULL COMMENT '资源的实例Id',
                            `vendor` tinyint(1) NOT NULL COMMENT '厂商\r\n',
                            `region` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
                            `create_at` bigint(20) NOT NULL,
                            `expire_at` bigint(20) NOT NULL,
                            `type` varchar(120) CHARACTER SET latin1 NOT NULL,
                            `name` varchar(255) NOT NULL,
                            `description` varchar(255) NOT NULL,
                            `status` varchar(255) CHARACTER SET latin1 NOT NULL,
                            `update_at` bigint(20) NOT NULL,
                            `sync_at` bigint(20) NOT NULL,
                            `accout` varchar(255) CHARACTER SET latin1 NOT NULL,
                            `public_ip` varchar(64) CHARACTER SET latin1 NOT NULL,
                            `private_ip` varchar(64) CHARACTER SET latin1 NOT NULL,
                            PRIMARY KEY (`id`),
                            KEY `name` (`name`) USING BTREE,
                            KEY `status` (`status`),
                            KEY `private_ip` (`public_ip`) USING BTREE,
                            KEY `public_ip` (`public_ip`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

CREATE TABLE `host` (
                        `resource_id` varchar(64) NOT NULL,
                        `cpu` tinyint(4) NOT NULL,
                        `memory` int(11) NOT NULL,
                        `gpu_amount` tinyint(4) DEFAULT NULL,
                        `gpu_spec` varchar(255) DEFAULT NULL,
                        `os_type` varchar(255) DEFAULT NULL,
                        `os_name` varchar(255) DEFAULT NULL,
                        `serial_number` varchar(120) DEFAULT NULL,
                        PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4