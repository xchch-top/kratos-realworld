drop table if exists user;
create table user
(
    id            bigint                  not null auto_increment primary key,
    username      varchar(50)  default '' not null comment '姓名',
    email         varchar(50)             not null comment '邮箱',
    bio           varchar(255) default '' not null comment 'bio',
    image         varchar(50)  default '' not null comment '头像链接',
    password_hash varchar(255)            not null comment '密码hash',
    created_at    datetime     default now() comment '创建时间',
    updated_at    datetime     default now() comment '更新时间',
    deleted_at    datetime     default null comment '删除时间'
);

drop table if exists article;
create table article
(
    id              bigint      not null auto_increment primary key,
    slug            varchar(50) not null comment '别名',
    title           varchar(50) not null comment '标题',
    description     varchar(255) default null comment '描述',
    body            varchar(50) not null comment '内容',
    author_id       bigint      not null comment '文章作者',
    favorites_count int          default 0 comment '点赞数',
    created_at      datetime     default now() comment '创建时间',
    updated_at      datetime     default now() comment '更新时间',
    deleted_at      datetime     default null comment '删除时间'
);