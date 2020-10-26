package service

import (
	"github.com/astaxie/beego"
	"gpm/models"
	"gpm/tools"
	"io/ioutil"
	"os"
)


type LocalFileSystem struct {
	RootPath string;
}

func (s *LocalFileSystem) ReadByte(parentDir string,fileName string) ([]byte,error){
	destPath:=s.RootPath+tools.PathSeparator+parentDir+tools.PathSeparator+fileName
	return ioutil.ReadFile(destPath)
}
func (s *LocalFileSystem) ReadText(parentDir string,fileName string) (string,error){
	readByte,err:=s.ReadByte(parentDir,fileName)
	return string(readByte),err;
}

func (s *LocalFileSystem) Mkdir(parentDir string,fileName string) error {
	destPath:=s.RootPath+tools.PathSeparator+parentDir+tools.PathSeparator+fileName
	//存在目录就跳过
	_, err := os.Stat(destPath)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
	}
	return os.Mkdir(destPath,os.ModePerm);
}

func (s *LocalFileSystem) RmDir(parentDir string,fileName string) error {
	destPath:=s.RootPath+tools.PathSeparator+parentDir+tools.PathSeparator+fileName
	return os.RemoveAll(destPath);
}
func (s *LocalFileSystem) ListRoot() ([]models.Node,error) {
	return s.ListDir(tools.PathSeparator)
}
func (s *LocalFileSystem) IsDir(destPath string) (bool) {
	destDirPath:=s.RootPath+tools.PathSeparator+destPath
	fi, _ := os.Stat(destDirPath)
	return fi.IsDir()
}
func (s *LocalFileSystem) ListDir(dirPth string) ([]models.Node,error) {
	destDirPath:=s.RootPath+tools.PathSeparator+dirPth
	dir, err := ioutil.ReadDir(destDirPath)
	nodeList:=make([]models.Node,len(dir));

	if err != nil {
		return nil, err
	}
	//PthSep := string(os.tools.PathSeparator)
	for index, fi := range dir {
		node:=models.Node{
			Title:fi.Name(),
			Expand:false,
			Contextmenu:true,
			IsDir:true,
			Children:[]models.Node{},
			DirPath:dirPth,
		}
		if fi.IsDir() {
			node.IsDir=true;
		}else{
			node.IsDir=false;
		}
		nodeList[index]=node;
	}
	return nodeList, nil
}


func (s *LocalFileSystem) DeleteFile(parentDir string,fileName string) error {
	destPath:=s.RootPath+tools.PathSeparator+parentDir+tools.PathSeparator+fileName
	return os.Remove(destPath);
}
func (s *LocalFileSystem) CreateFile(parentDir string,fileName string) error {
	destPath:=s.RootPath+tools.PathSeparator+parentDir+tools.PathSeparator+fileName
	f,err:=os.Create(destPath)
	defer f.Close()
	return err;
}
func (s *LocalFileSystem) SaveTextFile(parentDir string,fileName string,content string,policyType os.FileMode) error {
	destPath:=s.RootPath+tools.PathSeparator+parentDir+tools.PathSeparator+fileName
	return ioutil.WriteFile(destPath,[]byte(content),policyType)
}
func (s *LocalFileSystem) Rename(srcDir string,src string,dest string) error {
	srcPath:=s.RootPath+tools.PathSeparator+srcDir+tools.PathSeparator+src
	destPath:=s.RootPath+tools.PathSeparator+srcDir+tools.PathSeparator+dest
	return os.Rename(srcPath,destPath)
}
func (s *LocalFileSystem) Ping() error {
	return nil;
}
/**
---------------------------------------
 文件系统工厂类负责读取配置参数生成文件系统实例
---------------------------------------
 */
type LocalFileSystemFactory struct {
}
func (s *LocalFileSystemFactory) Create()  (FileSystem,error) {
	rootpath:=beego.AppConfig.String("rootpath")
	fileSystem:=LocalFileSystem{RootPath:rootpath}
	return &fileSystem,nil;
}
func (s *LocalFileSystemFactory) Name() string {
	return "service.LocalFileSystemFactory"
}
