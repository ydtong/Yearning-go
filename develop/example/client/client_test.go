package client

import (
	pb "Juno/libraGrpc/proto"
	"testing"
)

func TestExDDLClient(t *testing.T) {

	s := pb.LibraAuditOrder{
		SQL:      "CREATE TABLE `core_grai32ned` (`id` int unsigned AUTO_INCREMENT PRIMARY KEY COMMENT 'id',`username` varchar(50) NOT NULL COMMENT 'xxx')",
		DataBase: "test01",
		Source: &pb.Source{
			Addr:     "127.0.0.1",
			User:     "root",
			Password: "xxxxx",
			Port:     3306,
		},
		IsDML:   false,
		Backup:  true,
		Execute: true,
		WorkId:  "xxxxx123",
	}
	ExDDLClient(&s)
}

func TestExDMLClient(t *testing.T) {

	s := pb.LibraAuditOrder{
		SQL:      "INSERT INTO `Yearning`.`b_test`(`id`,`idx`,`dateux`) VALUES(66732326,'321',NULL);",
		DataBase: "test01",
		Table:    "b_test",
		Source: &pb.Source{
			Addr:     "127.0.0.1",
			User:     "root",
			Password: "xxxxx",
			Port:     3306,
		},
		IsDML:   true,
		Backup:  true,
		Execute: true,
		WorkId:  "xxxxx1234",
	}
	ExDMLClient(&s)
}

func TestTsClient(t *testing.T) {

	s := pb.LibraAuditOrder{
		SQL:      "CREATE TABLE `core_grai32ned` (`id` int unsigned AUTO_INCREMENT PRIMARY KEY COMMENT 'id',`username` varchar(50) NOT NULL COMMENT 'xxx')",
		DataBase: "test01",
		Source: &pb.Source{
			Addr:     "127.0.0.1",
			User:     "root",
			Password: "xxxxx",
			Port:     3306,
		},
		Check:  true,
		WorkId: "xxxxx123",
	}
	TsClient(&s)
}
