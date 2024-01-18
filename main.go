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
	Types []string
}

func main(){
	var err error
	var feats []string
	input_data := calculation.InputData{}
	
	var type_choice string
	var choice_selected bool = false
	for (choice_selected != true){
		fmt.Println("Add build type filtering?", " ", "(y/n)")
		fmt.Scanln(&type_choice)

		if type_choice == "y" || type_choice == "n" {
			choice_selected = true
		}
	}
	
	skill_limiter := 0
	stat_limiter := 35

	tmpl := generate_templates()
	
	err = take_stat_input(tmpl.Stats, &input_data, &stat_limiter)
	if err != nil {
		return
	}
	err = take_skill_input(tmpl.Skills, &input_data, &skill_limiter)
	if err != nil {
		return
	}
	print_values(input_data)
	var choice string
	choice_selected = false
	for (choice_selected != true){
		fmt.Println("Begin calculation?", " ", "(y/n)")
		fmt.Scanln(&choice)

		if choice == "y"{
			choice_selected = true
		}
	}

	feats, err = calculation.Prepare_Data(input_data)
	fmt.Println(feats)
	
	for f := 0; f < len(feats); f++ {
		input_data.Feats = append(input_data.Feats, feats[f])
	}
	
	feats_json, error := json.MarshalIndent(input_data, "", "    ")
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


func take_stat_input (stat_slice []string, input_data *calculation.InputData, stat_limiter *int) error {
	
	for i := 0; i < len(stat_slice); i++ {
		fmt.Println("-----------------------------------------------")
		fmt.Println()
		stat := stat_slice[i]
		choice_selected := false
		var val string
		var val_as_int int
		for (choice_selected != true){

			fmt.Println("Enter value for: ", stat)
			_, err := fmt.Scanln(&val)
			if err != nil {
				fmt.Println("Failed to read input..")
			}
		
			int_stat_val, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Integer conversion failed")
			}

			if err == nil {
				val_as_int = int_stat_val
				choice_selected = true
			}
		}
	
		if val_as_int > 20 {
			val_as_int = 20
		}

		if val_as_int < 3 {
			val_as_int = 3
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
	
		if val_as_int > 3 && val_as_int <= 20 && *stat_limiter <= 46{
			stat_data := calculation.Stat{
				StatName: stat,
				StatValue: val_as_int,
			}

			input_data.Stats = append(input_data.Stats, stat_data)
		} 


		if *stat_limiter <= 46 {
			fmt.Println("Stat point counter: ", *stat_limiter, "/46")	
		}

		if *stat_limiter > 46 {
			overlimit := (*stat_limiter - 46)
			val_as_int = (val_as_int - overlimit)
			*stat_limiter -= overlimit
			fmt.Println("Stat point counter: ", *stat_limiter, "/46")
			fmt.Println("All stat points used before completing allocations, assigning base value 5 to remaining stats.")
			stat_data := calculation.Stat {
				StatName: stat,
				StatValue: val_as_int,
			}
			input_data.Stats = append(input_data.Stats, stat_data)
			remainder := i+1
			assign_remaining_stats(remainder, stat_slice, input_data)
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			return nil
		}
		fmt.Println()
		fmt.Println("-----------------------------------------------")
	}
	return nil
}



func assign_remaining_stats(RI int, slice []string, input_data *calculation.InputData) {
	for r := RI; r < len(slice); r++ {
		stat := slice[r]
		stat_data := calculation.Stat {
			StatName: stat,
			StatValue: 5,
		}
		input_data.Stats = append(input_data.Stats, stat_data)
	}
}

func take_skill_input(skill_slice []string, input_data *calculation.InputData, skill_limiter *int) error {
	for i := 0; i < len(skill_slice); i++ {
		fmt.Println("-----------------------------------------------")
		fmt.Println()
		skill := skill_slice[i]
		choice_selected := false
		var val string
		var val_as_int int

		for(choice_selected != true){

			fmt.Println("Enter value for: ", skill)
			_, err := fmt.Scanln(&val)
			if err != nil {
				return err
			}
			
			int_skill_var, err := strconv.Atoi(val)
			if err != nil {
				return err
			}

			if err == nil {
				val_as_int = int_skill_var
				choice_selected = true
			}
		}
		fmt.Println("chosen value: ", skill, "->", val_as_int)
	
		if val_as_int > 160 {
			val_as_int = 160
		}

		if val_as_int > 0 {
			*skill_limiter += val_as_int
		}
	


		if val_as_int > 0 && val_as_int <= 160 && *skill_limiter <= 1280 {
			fmt.Println("Skill point counter: ", *skill_limiter, "/1280")
			skill_data := calculation.Skill{
				SkillName: skill,
				SkillValue: val_as_int,
			}
			input_data.Skills = append(input_data.Skills, skill_data)
		} 

		if *skill_limiter > 1280 {
			fmt.Println("All skill points used, assigning selected values")

			overlimit := (*skill_limiter - 1280)
			val_as_int = (val_as_int - overlimit)
			*skill_limiter -= overlimit

			skill_data := calculation.Skill {
				SkillName: skill,
				SkillValue: val_as_int,
			}
			input_data.Skills = append(input_data.Skills, skill_data)
			remainder := i+1
			fmt.Println("Skill point counter: ", *skill_limiter, "/1280")
			assign_remaining_skills(remainder, skill_slice, input_data)
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			return nil
		}
		fmt.Println()
		fmt.Println("-----------------------------------------------")
	}
	return nil
}

func assign_remaining_skills(RI int, slice []string, input_data *calculation.InputData){
	for r := RI; r < len(slice); r++ {
		skill := slice[r]
		skill_data := calculation.Skill {
			SkillName: skill,
			SkillValue: 0,
		}
		input_data.Skills = append(input_data.Skills, skill_data)
	}
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



func print_values (input_data calculation.InputData) {


	var concat_stat_item string
	var concat_skill_item string
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	for i := 0; i < len(input_data.Stats); i++ {
		item := input_data.Stats[i]
		item_num := strconv.Itoa(item.StatValue)
		concat_stat_item += item.StatName + "-->" + item_num + ", "
	}

	fmt.Println("STATS: ")
	fmt.Println(concat_stat_item)
	fmt.Println()

	for s := 0; s < len(input_data.Skills); s++ {
		item := input_data.Skills[s]
		item_num := strconv.Itoa(item.SkillValue)
		concat_skill_item += item.SkillName + "--> " + item_num + ", "
	}

	fmt.Println("SKILLS: ")
	fmt.Println(concat_skill_item)
	fmt.Println()
	fmt.Println("--------------------------------------------------")
}
