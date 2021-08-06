package pipe

import (
	"github.com/abc463774475/bbtool/n_log"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var pipeFile = "pipe.ipc"

func TestPipeServer(t *testing.T) {
	os.Remove(pipeFile)

	err := syscall.Mkfifo(pipeFile,0666)
	if err != nil {
		n_log.Panic("err  %v", err)
	}

	n_log.Info("start  ok")
	go func() {
		file, err := os.OpenFile(pipeFile, os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			n_log.Panic("err  %v",err)
		}

		for {
			data := make([]byte,1024)

			n,err := file.Read(data)
			if err != nil {
				n_log.Panic("err  %v", err)
			}

			n_log.Info("n  %v  msg\n%v", n, string(data[:n]))
		}
	}()

	time.Sleep(300*time.Second)
}

func TestPipeWrite(t *testing.T) {
	file, err := os.OpenFile(pipeFile, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		n_log.Panic("err  %v",err)
	}

	file.Write([]byte("ppppppp"))

	time.Sleep(10*time.Second)
}

func TestMap(t *testing.T)  {
	m := map[string]interface{}{
		"dddd":0,
	}


	n_log.Info("m  %v", m)

}


type st struct {
	ID  int	   `gorm:"id"`
	Val string `gorm:"val"`
	V1 uint8 `gorm:"v1"`
}

func (st ) TableName() string {
	return "t1"
}

type CoreQueryOrder struct {
	ID       uint            `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	WorkId   string          `gorm:"type:varchar(50);not null;index:workId_idx" json:"work_id"`
	// EndTime  string          `gorm:"index:idx_endTime,priority:1;type:char(16)" json:"end_time"` // 有效期
	// Username string          `gorm:"index:idx_endTime,priority:2;type:varchar(50);not null" json:"username"`
	EndTime  string          `gorm:"index:idx_endTime;type:char(16)" json:"end_time"` // 有效期
	Username string          `gorm:"index:idx_endTime;type:varchar(50);not null" json:"username"`
	Date     string          `gorm:"type:varchar(50);not null;index:query_idx" json:"date"`
	Text     string          `gorm:"type:longtext;not null" json:"text"`
	IDC      string          `gorm:"type:varchar(50);not null" json:"idc"`
	Source   string          `gorm:"type:varchar(50);" json:"source"`
	Assigned string          `gorm:"type:varchar(50);not null" json:"assigned"`
	Realname string          `gorm:"type:varchar(50);not null" json:"real_name"`
	Export   uint            `gorm:"type:tinyint(2);not null" json:"export"`
}

type User struct {
	Name   string `gorm:"index:idx_member"`
	Number string `gorm:"index:idx_member"`

	Val string `gorm:"index:idx_member"`
}

func TestGorm(t *testing.T)  {
	db,err :=gorm.Open("mysql","iot_rw:Zj1tmb^!Adm43geD@(172.17.2.69)/iot_resource?charset=utf8mb4&loc=Local")
	if err != nil {
		n_log.Panic("err   %v", err)
	}
	err = db.DB().Ping()
	n_log.Info("db  erro %v", err)
	//db = db.Raw("set max_execution_time = 2; ")
	//
	//
	////r := db.Exec("select /*+MAX_EXECUTION_TIME(20)*/  SLEEP(0.5) ,t1.id from t1;")
	//
	//n_log.Info("err \n%v",r.Error)
	//strs := &struct {
	//	SQL string `gorm:"sql"`
	//	Message string `gorm:"message"`
	//}{}
	//err = r.Scan(&strs).Error
	//
	//n_log.Info("err  %v\n%v", err, strs)
}