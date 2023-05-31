package main

import (
	"fmt"
	"web/core"
	_ "web/docs"
	"web/flag"
	"web/global"
	"web/routers"
)

// @title web API文档
// @version 1.0
// @description web API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	core.InitConf()
	fmt.Println(global.Config)
	global.DB = core.InitGorm()
	// fmt.Println(global.DB)
	global.Log = core.InitLogger()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 连接redis
	global.Redis = core.ConnectRedis()
	// 连接es
	global.ESClient = core.EsConnect()

	router := routers.InitRouter()
	err := router.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}



// 以下代码测试es用
// 使用前需要把以上代码注释
// import(
// 	"github.com/olivere/elastic/v7"
// 	"github.com/sirupsen/logrus"
// 	"web/core"
// 	"context"
// )
// var client *elastic.Client

// func EsConnect() *elastic.Client  {
//   var err error
//   sniffOpt := elastic.SetSniff(false)
//   host := "http://127.0.0.1:9200"
//   c, err := elastic.NewClient(
//     elastic.SetURL(host),
//     sniffOpt,
//     elastic.SetBasicAuth("", ""),
//   )
//   if err != nil {
//     logrus.Fatalf("es连接失败 %s", err.Error())
//   }
//   return c
// }

// func init() {
// 	core.InitConf()
// 	core.InitLogger()
//   client = EsConnect()

// }

// type DemoModel struct {
//   ID        string `json:"id"`
//   Title     string `json:"title"`
//   UserID    uint   `json:"user_id"`
//   CreatedAt string `json:"created_at"`
// }

// func (DemoModel) Index() string {
//   return "demo_index"
// }

// func (DemoModel) Mapping() string {
//   return `
// {
//   "settings": {
//     "index":{
//       "max_result_window": "100000"
//     }
//   }, 
//   "mappings": {
//     "properties": {
//       "title": { 
//         "type": "text"
//       },
//       "user_id": {
//         "type": "integer"
//       },
//       "created_at":{
//         "type": "date",
//         "null_value": "null",
//         "format": "[yyyy-MM-dd HH:mm:ss]"
//       }
//     }
//   }
// }
// `
// }

// // IndexExists 索引是否存在
// func (demo DemoModel) IndexExists() bool {
//   exists, err := client.
//     IndexExists(demo.Index()).
//     Do(context.Background())
//   if err != nil {
//     logrus.Error(err.Error())
//     return exists
//   }
//   return exists
// }

// // CreateIndex 创建索引
// func (demo DemoModel) CreateIndex() error {
//   if demo.IndexExists() {
//     // 有索引
//     demo.RemoveIndex()
//   }
//   // 没有索引
//   // 创建索引
//   createIndex, err := client.
//     CreateIndex(demo.Index()).
//     BodyString(demo.Mapping()).
//     Do(context.Background())
//   if err != nil {
//     logrus.Error("创建索引失败")
//     logrus.Error(err.Error())
//     return err
//   }
//   if !createIndex.Acknowledged {
//     logrus.Error("创建失败")
//     return err
//   }
//   logrus.Infof("索引 %s 创建成功", demo.Index())
//   return nil
// }

// // RemoveIndex 删除索引
// func (demo DemoModel) RemoveIndex() error {
//   logrus.Info("索引存在，删除索引")
//   // 删除索引
//   indexDelete, err := client.DeleteIndex(demo.Index()).Do(context.Background())
//   if err != nil {
//     logrus.Error("删除索引失败")
//     logrus.Error(err.Error())
//     return err
//   }
//   if !indexDelete.Acknowledged {
//     logrus.Error("删除索引失败")
//     return err
//   }
//   logrus.Info("索引删除成功")
//   return nil
// }

// func main() {
// 	DemoModel{}.CreateIndex()
// }
