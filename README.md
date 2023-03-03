# ddd-go-webapp-sample

## Table DDL
```sql
CREATE TABLE `user` (
`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'user id',
`account` varchar(20) NOT NULL COMMENT 'account',
`password` varchar(20) NOT NULL COMMENT 'password',
`create_time` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'create_time',
`update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'update_time',
PRIMARY KEY (`id`),
UNIQUE KEY `UK_user_acc` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='user';
```