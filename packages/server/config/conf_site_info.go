package config

type SiteInfo struct {
  CreatedAt string `yaml:"created_at" json:"created_at"`
  BeiAn 		string `yaml:"bei_an" 	  json:"bei_an"`
  Title  		string `yaml:"title" 		  json:"title"`
	QQImage 	string `yaml:"qq_image" 	json:"qq_image"`
  Version 	string `yaml:"version" 	  json:"version"`
  Email  		string `yaml:"email" 	    json:"email"`
	// wechat_image string `yaml: "host"`
  // name int `yaml: "port"`
  // job  string `yaml: "env"`
	// addr string `yaml: "host"`
  // slogan int `yaml: "port"`
  // slogan_en  string `yaml: "env"`
	// web string `yaml: "host"`
  
}
