package g

import (
	"github.com/Sirupsen/logrus"
	"github.com/kless/osutil/user/crypt/sha512_crypt"
	"jadegong/api.mmsystem.com/model"
)

func initRoot() {
	if len(Conf.RootEmail) == 0 || len(Conf.RootPassword) == 0 {
		logrus.Error("Please config root user email and password!")
		return
	}
	if IsEmail(Conf.RootEmail) == false {
		logrus.Error("Invalid root user email format!")
		return
	}

	user := &model.User{}
	session := Session()
	db := session.DB(Conf.DBName)
	defer session.Close()
	//TODO add root user

	c := sha512_crypt.New()
	rootPassword, _ := c.Generate([]byte(Conf.RootPassword), []byte(""))
}
