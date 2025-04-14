# GO API TEMPLATE
![alt text](image.png)

## Tools used:

*  sqlc **[CLI]**
*  golang-migrate **[CLI]**
*  viper
*  net/http (without framework) 


## Architecture (clean based)


📦 App (shared) <br>
├── 📦cmd<br>
│   └── _main_files_<br>


├── 📦configs<br>
│   └── _config_files_<br>


🏗️ Internal<br>
├── 📦 Presentation<br>
├── 📦 Use Cases<br>
│   └── 📄 Queries <br>
├── 📦 Repositories <br>
│   └── 📄 Queries <br>

🗄️ SQL Layer <br>
├── 📤 Migrations <br>
└── 📄 Queries
