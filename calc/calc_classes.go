package calculation

type Stat struct {
	StatName  string
	StatValue int   
}

type Skill struct {
	SkillName  string
	SkillValue int    
}

type InputData struct {
	Stats          []Stat 
	Skills         []Skill 
	Character_Type []string 
}

type Skills_Tracker struct {
	Skill_Name  string
	Skill_Value int
	Feat_Value  string
}

type Stats_Tracker struct {
	Stat_Name  string
	Stat_Value int
	Feat_Value string
}








