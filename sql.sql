CREATE DATABASE `vedio` /*!40100 DEFAULT CHARACTER SET utf8 */;
CREATE TABLE `comments` (
  `id` varchar(64) NOT NULL,
  `vedio_id` varchar(64) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  `content` text,
  `time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `sessions` (
  `session_id` varchar(255) NOT NULL,
  `TTL` tinytext,
  `login_name` varchar(64) DEFAULT NULL COMMENT '用户名',
  PRIMARY KEY (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login_name` varchar(64) NOT NULL COMMENT '登录用户名',
  `pwd` varchar(255) NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `vedio_info` (
  `id` varchar(64) NOT NULL,
  `author_id` int(11) DEFAULT NULL COMMENT '用户id',
  `name` text COMMENT '标题',
  `display_ctime` text COMMENT '显示时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
