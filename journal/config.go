package journal

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"runtime"
)

const (
	CONFIG = "config.json"
)

type Config struct {
	DateFormatter string `mapstructure:"DATE_FORMAT"`
	DBType        string `mapstructure:"DBTYPE"`
	OSType        string
	DBPath        string
	DataPath      string
	ConfigPath    string
}

const (
	InterJournalDir = "interjournal"
	ConfigFile      = "config.json"
)

// NewConfig builds the data and config paths and returns a pointer to the config
func NewConfig() *Config {
	dbPath := GetDataPath() + "/journal.db"
	return &Config{
		DateFormatter: "",
		DBType:        "normal",
		OSType:        runtime.GOOS,
		DBPath:        dbPath,
		DataPath:      GetDataPath(),
		ConfigPath:    GetConfigPath(),
	}
}

func GetConfigPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dirname += "/.config/interjournal/"

	return dirname
}

//CreateConfigFile creates the config file and
func (j *Journal) CreateConfigFile() error {
	err := os.MkdirAll(j.Config.ConfigPath, 0755)
	if err != nil {
		return err
	}
	cfgPath := GetConfigPath()
	createEmptyFile(cfgPath + CONFIG)
	return nil
}

//createEmptyFile creates the config.json file in host config location
func createEmptyFile(name string) {
	d := []byte("")
	err := os.WriteFile(name, d, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// GetDataPath fetches the location of Interjournal's storage
// This will be the default location of the database
func GetDataPath() string {
	switch runtime.GOOS {
	case "darwin":
		cd, err := os.UserConfigDir()
		if err != nil {
			log.Fatal("config directory not found", err)
		}
		return cd + "/interjournal"
	}
	return ""
}

func (c *Config) createDataPath() {
	_, err := os.Stat(c.DataPath)
	if err != nil {
		log.Printf("interjournal data path missing\n error: %v", err)
		log.Println("creating interjournal data path: %v", c.DataPath)
		err := os.MkdirAll(c.DataPath, 0655)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Setup creates missing directories and files
// flag --setup <message>
func (j *Journal) Setup() error {
	// make the application directory
	err := j.CreateAppData()
	if err != nil {
		return err
	}
	// make the configuration file
	err = j.CreateConfigFile()
	if err != nil {
		return err
	}
	return nil
}
func (j *Journal) LoadConfig(path string) (config Config, err error) {
	// Check for correct dir
	// log missing config
	// create file
	// https://hub.jmonkeyengine.org/t/appdata-equivalent-on-macos-and-linux/43735/2
	// create datapath

	// Load Config
	// Set a default location for db
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// The installation directory is platform-specific:
//
//On Linux, the default is /opt/application-name
//On macOS, the default is /Applications/application-name
//On Windows, the default is c:\Program Files\application-name; if the --win-per-user-install option is used, the default is C:\Users\user-name\AppData\Local\application-name
// GetAppData will find the default location for your applicaiton to
func (j *Journal) CreateAppData() error {
	d, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	j.ApplicationDir = d + "/interjournal"

	err = os.MkdirAll(j.ApplicationDir, 0755)
	if err != nil {
		return err
	}
	return nil
}
