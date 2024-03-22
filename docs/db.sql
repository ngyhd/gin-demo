create table `gin_demo`.user
(
    id         int auto_increment
        primary key,
    username   varchar(20) not null comment '用户名',
    password   varchar(50) not null comment '账户密码',
    email      varchar(50) not null comment '邮箱',
    created_at int         not null comment '创建时间',
    updated_at int         not null comment '更新时间',
    deleted_at int         not null comment '删除时间'
)
    comment '用户表';

