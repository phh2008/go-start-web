/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : go_start_web

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 23/10/2022 18:47:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (10, 'p', 'admin', '/api/v1/addAuth', 'GET', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (12, 'p', 'admin', '/api/v1/delAuth', 'GET', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (13, 'p', 'admin', '/api/v1/hello', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', 'admin', '/api/v3/hello', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', 'admin2', '/api/v1/hello2', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', 'admin3', '/api/v3/hello', 'POST', '', '', '', '', '');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `amount` decimal(20, 2) NULL DEFAULT NULL,
  `create_at` datetime(0) NULL DEFAULT NULL,
  `status` smallint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (2695795661258752, 1, '广东aa有限公司', 550.00, '2018-05-14 00:00:00', 2);
INSERT INTO `order` VALUES (2695800354586624, 1, '广东aa有限公司', 1210.00, '2018-05-14 00:00:00', 5);
INSERT INTO `order` VALUES (2696688667082752, 1, '东莞cc公司', 47.80, '2018-05-15 09:09:15', 5);
INSERT INTO `order` VALUES (2696755529547776, 1, '广东aa有限公司', 13.00, '2018-05-15 00:00:00', 5);
INSERT INTO `order` VALUES (2696763584970752, 1, '东莞cc公司', 16.42, '2018-05-15 10:30:40', 5);
INSERT INTO `order` VALUES (2696767262343168, 1, '广东aa有限公司', 180.00, '2018-05-15 10:32:37', 5);
INSERT INTO `order` VALUES (2696775042531328, 1, '广东aa有限公司', 100.00, '2018-05-15 10:38:35', 5);
INSERT INTO `order` VALUES (2696790954049536, 1, '东cc公司', 11.75, '2018-05-15 10:53:03', 5);
INSERT INTO `order` VALUES (2696792718098432, 1, '广东aa有限公司', 20.00, '2022-05-04 15:04:59', 5);
INSERT INTO `order` VALUES (2696797306683392, 2, '东莞cc公司', 2.60, '2018-05-15 10:59:29', 4);
INSERT INTO `order` VALUES (2696800312573952, 2, '东莞cc公司', 4890.00, '2018-05-15 00:00:00', 5);
INSERT INTO `order` VALUES (2696839197065216, 2, '东莞cc公司', 117.50, '2018-05-15 11:41:15', 5);
INSERT INTO `order` VALUES (2696992436191232, 2, '东莞cc公司', 2.39, '2018-05-15 14:18:16', 5);
INSERT INTO `order` VALUES (2696996734763008, 2, '东莞cc公司', 2.39, '2018-05-15 14:22:38', 5);
INSERT INTO `order` VALUES (2697019514847232, 2, '广东aa有限公司', 300.00, '2018-05-15 14:47:18', 6);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `age` int(11) NULL DEFAULT NULL,
  `passwd` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `birthday` datetime(0) NULL DEFAULT NULL,
  `created` datetime(0) NULL DEFAULT NULL,
  `updated` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'jinzhu', '', 0, '123456', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `user` VALUES (2, 'jack', 'abc', 24, '123456', '2001-06-14 10:34:41', '2022-03-27 19:44:41', '2022-03-27 19:44:41');

SET FOREIGN_KEY_CHECKS = 1;
