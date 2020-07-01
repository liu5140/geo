

CREATE TABLE `geo_table` (
`id`  bigint(20) NOT NULL AUTO_INCREMENT ,
`created_at`  timestamp NULL DEFAULT NULL ,
`updated_at`  timestamp NULL DEFAULT NULL ,
`deleted_at`  timestamp NULL DEFAULT NULL ,
`table_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`table_comment`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`class_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`tpl_category`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`package_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`module_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`business_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`function_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`function_author`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`options`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`create_by`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`update_by`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`remark`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=55
ROW_FORMAT=COMPACT
;



CREATE TABLE `geo_table_columns` (
`id`  bigint(20) NOT NULL AUTO_INCREMENT ,
`created_at`  timestamp NULL DEFAULT NULL ,
`updated_at`  timestamp NULL DEFAULT NULL ,
`deleted_at`  timestamp NULL DEFAULT NULL ,
`column_id`  bigint(20) NULL DEFAULT NULL ,
`table_id`  bigint(20) NULL DEFAULT NULL ,
`column_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`column_comment`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`column_type`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`go_type`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`go_field`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`html_field`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`is_pk`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`is_increment`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`is_required`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`is_insert`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`is_edit`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`is_list`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`is_query`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`query_type`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`html_type`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`dict_type`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`sort`  int(11) NULL DEFAULT NULL ,
`create_by`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`update_by`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=584
ROW_FORMAT=COMPACT
;



INSERT INTO `ds`.`geo_table` (`id`, `created_at`, `updated_at`, `deleted_at`, `table_name`, `table_comment`, `class_name`, `tpl_category`, `package_name`, `module_name`, `business_name`, `function_name`, `function_author`, `options`, `create_by`, `update_by`, `remark`) VALUES ('53', '2020-03-12 18:05:41', NULL, NULL, 't_answer', '回复表', 'answer', 'crud', 'yj-app', 'model', 'Answer', '回复', 'yunjie', '', 'admin', '', '');
INSERT INTO `ds`.`geo_table` (`id`, `created_at`, `updated_at`, `deleted_at`, `table_name`, `table_comment`, `class_name`, `tpl_category`, `package_name`, `module_name`, `business_name`, `function_name`, `function_author`, `options`, `create_by`, `update_by`, `remark`) VALUES ('54', '2020-06-28 11:51:30', NULL, NULL, 'sys_dept', '部门表', 'dept', 'crud', 'yj-app', 'model', 'Dept', '部门', 'yunjie', '', 'admin', '', '');



INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('557', NULL, NULL, NULL, NULL, '53', 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', '1', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('558', NULL, NULL, NULL, NULL, '53', 'pid', '问题ID', 'bigint(20)', 'int64', 'Pid', 'pid', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '2', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('559', NULL, NULL, NULL, NULL, '53', 'atype', '回复人类别', 'tinyint(1)', 'int', 'Atype', 'atype', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', '3', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('560', NULL, NULL, NULL, NULL, '53', 'user_id', '回复人ID', 'bigint(20)', 'int64', 'UserId', 'userId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '4', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('561', NULL, NULL, NULL, NULL, '53', 'nick_name', '回复人名称', 'varchar(50)', 'string', 'NickName', 'nickName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', '5', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('562', NULL, NULL, NULL, NULL, '53', 'avatar', '回复人头像', 'varchar(50)', 'string', 'Avatar', 'avatar', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '6', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('563', NULL, NULL, NULL, NULL, '53', 'remark', '回复内容', 'tinytext', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', '7', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('564', NULL, NULL, NULL, NULL, '53', 'img1', '回复图片1', 'varchar(100)', 'string', 'Img1', 'img1', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '8', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('565', NULL, NULL, NULL, NULL, '53', 'img2', '回复图片2', 'varchar(100)', 'string', 'Img2', 'img2', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '9', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('566', NULL, NULL, NULL, NULL, '53', 'img3', '回复图片3', 'varchar(100)', 'string', 'Img3', 'img3', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '10', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('567', NULL, NULL, NULL, NULL, '53', 'status', '状态', 'tinyint(1)', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', '11', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('568', NULL, NULL, NULL, NULL, '53', 'created_at', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', '12', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('569', NULL, NULL, NULL, NULL, '53', 'updated_at', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', '13', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('570', NULL, NULL, NULL, NULL, '54', 'dept_id', '部门id', 'bigint(20)', 'int64', 'DeptId', 'deptId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', '1', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('571', NULL, NULL, NULL, NULL, '54', 'parent_id', '父部门id', 'bigint(20)', 'int64', 'ParentId', 'parentId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '2', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('572', NULL, NULL, NULL, NULL, '54', 'ancestors', '祖级列表', 'varchar(50)', 'string', 'Ancestors', 'ancestors', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '3', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('573', NULL, NULL, NULL, NULL, '54', 'dept_name', '部门名称', 'varchar(30)', 'string', 'DeptName', 'deptName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', '4', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('574', NULL, NULL, NULL, NULL, '54', 'order_num', '显示顺序', 'int(4)', 'int', 'OrderNum', 'orderNum', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '5', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('575', NULL, NULL, NULL, NULL, '54', 'leader', '负责人', 'varchar(20)', 'string', 'Leader', 'leader', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '6', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('576', NULL, NULL, NULL, NULL, '54', 'phone', '联系电话', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '7', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('577', NULL, NULL, NULL, NULL, '54', 'email', '邮箱', 'varchar(50)', 'string', 'Email', 'email', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', '8', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('578', NULL, NULL, NULL, NULL, '54', 'status', '部门状态（0正常 1停用）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', '9', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('579', NULL, NULL, NULL, NULL, '54', 'del_flag', '删除标志（0代表存在 2代表删除）', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', '10', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('580', NULL, NULL, NULL, NULL, '54', 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', '11', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('581', NULL, NULL, NULL, NULL, '54', 'created_at', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', '12', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('582', NULL, NULL, NULL, NULL, '54', 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', '13', 'admin', '');
INSERT INTO `ds`.`geo_table_columns` (`id`, `created_at`, `updated_at`, `deleted_at`, `column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `update_by`) VALUES ('583', NULL, NULL, NULL, NULL, '54', 'updated_at', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', '14', 'admin', '');
