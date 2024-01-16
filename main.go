package main

import (
	"encoding/json"
	"fmt"
	"main/calc"
	"os"
	"strconv"
)

type Templates struct {
	Stats []string
	Skills []string
}

func main(){
	var err error
	var output string
	var feats []string
	input_data := calculation.InputData{}
	
	skill_limiter := 0
	stat_limiter := 35

	tmpl := generate_templates()
	
	err, output = take_stat_input(tmpl.Stats, &input_data, &stat_limiter)
	if err != nil || output == "Incomplete" {
		return
	}
	err, output = take_skill_input(tmpl.Skills, &input_data, &skill_limiter)
	if err != nil || output == "Incomplete" {
		return
	}

	feats, err = calculation.Prepare_Data(input_data)
	fmt.Println(feats)

	feats_json, error := json.MarshalIndent(feats, "", "    ")
	if error != nil{
		return
	}

	path := "builds/"
	var filename string
	fmt.Println("Enter file name")
	fmt.Scanln(&filename)

	name_and_path := path + filename + ".json"

	err = os.WriteFile(name_and_path, feats_json, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}
}


func take_stat_input (stat_slice []string, input_data *calculation.InputData, stat_limiter *int) (error, string) {
	
	for i := 0; i < len(stat_slice); i++ {
		stat := stat_slice[i]
		var val string
		fmt.Println("Enter value for: ", stat)
		_, err := fmt.Scanln(&val)
		if err != nil {
			return err, ""
		}
		
		val_as_int, err := strconv.Atoi(val)
		if err != nil {
			return err, ""
		}
	
		if val_as_int > 20 {
			val_as_int = 20
		}
		fmt.Println("chosen val: ", stat, "->" ,val_as_int)

		if val_as_int < 5 && val_as_int >= 3 {
			*stat_limiter -= (5 - val_as_int) 
		} else if val_as_int > 5 && val_as_int <= 20 {	
			*stat_limiter += (val_as_int - 5)
		}

		if val_as_int == 5 {
			*stat_limiter += 0
		}


		fmt.Println("Stat point counter: ", *stat_limiter, "/46")
	
		if val_as_int > 3 && val_as_int <= 20 && *stat_limiter < 46{
			stat_data := calculation.Stat{
				StatName: stat,
				StatValue: val_as_int,
			}

			input_data.Stats = append(input_data.Stats, stat_data)
		} 
		
		if *stat_limiter > 46 {
			fmt.Println("All stat points used before completing allocations, try again..")
			return nil,"Incomplete"
		}
	} 

	return nil, ""
}

func take_skill_input(skill_slice []string, input_data *calculation.InputData, skill_limiter *int) (error, string) {
	for i := 0; i < len(skill_slice); i++ {
		skill := skill_slice[i]
		var val string
		fmt.Println("Enter value for: ", skill)
		_, err := fmt.Scanln(&val)
		if err != nil {
			return err, ""
		}
		
		val_as_int, err := strconv.Atoi(val)
		if err != nil {
			return err, ""
		}
		fmt.Println("chosen value: ", skill, "->", val_as_int)
	
		if val_as_int > 160 {
			val_as_int = 160
		}

		if val_as_int > 0 {
			*skill_limiter += val_as_int
		}
	
		fmt.Println("Skill point counter: ", *skill_limiter, "/1280")

		if val_as_int > 0 && val_as_int <= 160 && *skill_limiter < 1280 {
			skill_data := calculation.Skill{
				SkillName: skill,
				SkillValue: val_as_int,
			}
			input_data.Skills = append(input_data.Skills, skill_data)
		} 

		if *skill_limiter > 1280 {
			fmt.Println("All skill points used, assigning selected values")
			return nil, ""
		}
	}
	return nil, ""
}

func generate_templates() Templates {

	var SkillSlice []string = []string{
		"Guns", "Throwing", "Crossbows", "Melee", "Dodge", "Evasion", "Stealth",
		"Hacking", "Lockpicking", "Pickpocketing", "Traps",
		"Mechanics", "Temporal Manipulation", "Persuasion",
		"Intimidation", "Mercantile", "Metathermics", "Psychokinesis",
		"Thought Control", "Tailoring", "Biology", "Chemistry", "Electronics",
	}
	var StatSlice []string = []string{"Strength", "Dexterity", "Agility", "Constitution", "Perception", "Will", "Intelligence"}


	tmpl := Templates {
		Stats: StatSlice,
		Skills: SkillSlice,
	}

	return tmpl
}
