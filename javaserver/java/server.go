package java

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/codeskyblue/go-sh"
	"github.com/go-cmd/cmd"
)

// const javaDir string = "/data/service/"
const spring_profiles_active string = "main2"

func Start(serviceName string) {
	fmt.Printf("#\t start\t %v\t   \n", serviceName)
	fmt.Printf("%v 正在启动 \n", serviceName)
	session := sh.NewSession()
	cmdstring := fmt.Sprintf("nohup java  -Xms1024m -Xmx3024m  -jar    -Dspring.profiles.active=%v  /data/service/%v/*.jar   >/data/service/logs/%v.log 2>&1 &", spring_profiles_active, serviceName, serviceName)
	fmt.Println(cmdstring)
	err := session.Command("sh", "-c", cmdstring).Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result, _ := CkeckStatus(serviceName)
	if result {
		fmt.Println("启动成功")
	}

}

func Stop(serviceName string) bool {
	fmt.Printf("#\t stop\t %v\t  \n", serviceName)
	b, pid := CkeckStatus(serviceName)
	if !b { // = false
		return true //服务没有运行
	}
	fmt.Printf("%v  pid: %v \n", serviceName, pid.(string))
	err := KillProcess(pid.(string))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true //成功 kill

}

func Restart(serviceNmae string) {
	result := Stop(serviceNmae)
	time.Sleep(time.Second * 1)
	if result {
		Start(serviceNmae)
	}
}

func CkeckStatus(serviceName string) (bool, interface{}) {
	cmdstring := fmt.Sprintf("jps  | grep %v  | awk '{print $1}'", serviceName)
	//fmt.Println(cmdstring)
	c := cmd.NewCmd("bash", "-c", cmdstring)
	<-c.Start()
	//fmt.Println(c.Status().Stdout)
	if len(c.Status().Stdout) < 1 {
		fmt.Printf("* \t %v\t 服务没有运行\n", serviceName)
		return false, nil
	}
	//检测到 服务运行返回 true and pid
	return true, c.Status().Stdout[0]
}

func KillProcess(pid string) error {
	// 执行命令并结束进程
	cmd := exec.Command("kill", "-9", fmt.Sprintf("%v", pid))
	err := cmd.Run()
	if err != nil {
		return err
	}
	//成功 kill 返回 nil
	return nil
}
