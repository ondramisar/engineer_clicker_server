package main


type DefaultMachine struct {
	ID        string   `firestore:"id,omitempty"`
	Name      string   `firestore:"name,omitempty"`
	TimeToReach  int   `firestore:"timeToReach,omitempty"`
	Cost		 int   `firestore:"cost,omitempty"`
	NumberOfMaterialsToGive		 int   `firestore:"numberOfMaterialsToGive,omitempty"`
	IdMaterialToGive	string	`firestore:"idMaterialToGive,omitempty"`
}

type UserMachine struct {
	ID        string   `firestore:"id,omitempty"`
	Name      string   `firestore:"name,omitempty"`
	TimeToReach  int64   `firestore:"timeToReach,omitempty"`
	NumberOfMaterialsToGive		 int   `firestore:"numberOfMaterialsToGive,omitempty"`
	IdMaterialToGive		string	`firestore:"idMaterialToGive,omitempty"`
	Lvl			int	`firestore:"lvl,omitempty"`
	WorkerId	string	`firestore:"workerId,omitempty"`
	IdUser		string `firestore:"idUser,omitempty"`
}

type DefaultWorker struct {
	ID        string   `firestore:"id,omitempty"`
	Name      string   `firestore:"name,omitempty"`
	TimeCutBy  				float64   `firestore:"timeCutBy,omitempty"`
	MaterialMultiplayer		int   	`firestore:"materialMultiplayer,omitempty"`
	Payment					int		`firestore:"payment,omitempty"`
}

type UserWorker struct {
	ID        string   `firestore:"id,omitempty"`
	Name      string   `firestore:"name,omitempty"`
	TimeCutBy  					float64   `firestore:"timeCutBy,omitempty"`
	MaterialMultiplayer		int   	`firestore:"materialMultiplayer,omitempty"`
	Payment					int		`firestore:"payment,omitempty"`
	Lvl			int	`firestore:"lvl,omitempty"`
	IsOnMachine			bool	`firestore:"isOnMachine,omitempty"`
	IdUser		string `firestore:"idUser,omitempty"`
}

type DefaultMaterial struct{
	ID			string 		 `firestore:"id,omitempty"`
	Value		int				 `firestore:"value,omitempty"`
	Name 		string			 `firestore:"name,omitempty"`
}

type UserMaterial struct{
	ID			string 		 `firestore:"id,omitempty"`
	Value		int				 `firestore:"value,omitempty"`
	Name 		string			 `firestore:"name,omitempty"`
	NumberOf		int64				 `firestore:"numberOf,omitempty"`
	DefaultMaterialId 		string 	`firestore:"defaultMaterialId,omitempty"`
	IdUser		string `firestore:"idUser,omitempty"`
}

type User struct{
	IdUser 	string	`firestore:"idUser,omitempty"`
	Name	string	`firestore:"name,omitempty"`
	Email	string	`firestore:"email,omitempty"`
	Coins	 int	`firestore:"coins,omitempty"`
	LastUpdateMaterial int	`firestore:"lastUpdateMaterial,omitempty"`
	LastTimeOutOfApp	int64	`firestore:"lastTimeOutOfApp,omitempty"`
	LastPayment	int64	`firestore:"lastPayment,omitempty"`
}