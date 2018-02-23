package genomelink

type (
	// PhenotypeName is URL-safe phenotype name.
	PhenotypeName string

	// Phenotype represents phenotype definition that includes display name and category.
	Phenotype struct {
		URLName     PhenotypeName `json:"url_name"`
		DisplayName string        `json:"display_name"`
		Category    string        `json:"category"`
	}
)

const (
	PhenotypeNameAgreeableness                             PhenotypeName = "agreeableness"
	PhenotypeNameAlcoholDrinkingBehavior                   PhenotypeName = "alcohol-drinking-behavior"
	PhenotypeNameAlphaLinolenicAcid                        PhenotypeName = "alpha-linolenic-acid"
	PhenotypeNameAnger                                     PhenotypeName = "anger"
	PhenotypeNameBeardThickness                            PhenotypeName = "beard-thickness"
	PhenotypeNameBetaCarotene                              PhenotypeName = "beta-carotene"
	PhenotypeNameBitterTaste                               PhenotypeName = "bitter-taste"
	PhenotypeNameBlackHair                                 PhenotypeName = "black-hair"
	PhenotypeNameBloodGlucose                              PhenotypeName = "blood-glucose"
	PhenotypeNameBMI                                       PhenotypeName = "bmi"
	PhenotypeNameBodyFatMass                               PhenotypeName = "body-fat-mass"
	PhenotypeNameBodyFatPercentage                         PhenotypeName = "body-fat-percentage"
	PhenotypeNameBreastSize                                PhenotypeName = "breast-size"
	PhenotypeNameCaffeineConsumption                       PhenotypeName = "caffeine-consumption"
	PhenotypeNameCaffeineMetaboliteRatio                   PhenotypeName = "caffeine-metabolite-ratio"
	PhenotypeNameCalcium                                   PhenotypeName = "calcium"
	PhenotypeNameCarbohydrateIntake                        PhenotypeName = "carbohydrate-intake"
	PhenotypeNameChildhoodIntelligence                     PhenotypeName = "childhood-intelligence"
	PhenotypeNameConscientiousness                         PhenotypeName = "conscientiousness"
	PhenotypeNameDepression                                PhenotypeName = "depression"
	PhenotypeNameEggAllergy                                PhenotypeName = "egg-allergy"
	PhenotypeNameEndurancePerformance                      PhenotypeName = "endurance-performance"
	PhenotypeNameExcessiveDaytimeSleepiness                PhenotypeName = "excessive-daytime-sleepiness"
	PhenotypeNameExtraversion                              PhenotypeName = "extraversion"
	PhenotypeNameGeneticEyeColor                           PhenotypeName = "eye-color"
	PhenotypeNameFolate                                    PhenotypeName = "folate"
	PhenotypeNameFreckles                                  PhenotypeName = "freckles"
	PhenotypeNameGambling                                  PhenotypeName = "gambling"
	PhenotypeNameHarmAvoidance                             PhenotypeName = "harm-avoidance"
	PhenotypeNameHearingFunction                           PhenotypeName = "hearing-function"
	PhenotypeNameHeight                                    PhenotypeName = "height"
	PhenotypeNameHippocampalVolume                         PhenotypeName = "hippocampal-volume"
	PhenotypeNameIntelligence                              PhenotypeName = "intelligence"
	PhenotypeNameIron                                      PhenotypeName = "iron"
	PhenotypeNameJobRelatedExhaustion                      PhenotypeName = "job-related-exhaustion"
	PhenotypeNameLeanBodyMass                              PhenotypeName = "lean-body-mass"
	PhenotypeNameLobeSize                                  PhenotypeName = "lobe-size"
	PhenotypeNameLongevity                                 PhenotypeName = "longevity"
	PhenotypeNameMagnesium                                 PhenotypeName = "magnesium"
	PhenotypeNameMalePatternBaldnessAGA                    PhenotypeName = "male-pattern-baldness-aga"
	PhenotypeNameMathematicalAbility                       PhenotypeName = "mathematical-ability"
	PhenotypeNameMilkAllergy                               PhenotypeName = "milk-allergy"
	PhenotypeNameMorningPerson                             PhenotypeName = "morning-person"
	PhenotypeNameMotionSickness                            PhenotypeName = "motion-sickness"
	PhenotypeNameNeuroticism                               PhenotypeName = "neuroticism"
	PhenotypeNameNoveltySeeking                            PhenotypeName = "novelty-seeking"
	PhenotypeNameOpenness                                  PhenotypeName = "openness"
	PhenotypeNamePeanutsAllergy                            PhenotypeName = "peanuts-allergy"
	PhenotypeNamePhosphorus                                PhenotypeName = "phosphorus"
	PhenotypeNameProteinIntake                             PhenotypeName = "protein-intake"
	PhenotypeNameReadingAndSpellingAbility                 PhenotypeName = "reading-and-spelling-ability"
	PhenotypeNameRedHair                                   PhenotypeName = "red-hair"
	PhenotypeNameRedWineLiking                             PhenotypeName = "red-wine-liking"
	PhenotypeNameResponseToVitaminESupplementation         PhenotypeName = "response-to-vitamin-e-supplementation"
	PhenotypeNameRewardDependence                          PhenotypeName = "reward-dependence"
	PhenotypeNameSkinPigmentation                          PhenotypeName = "skin-pigmentation"
	PhenotypeNameSleepDuration                             PhenotypeName = "sleep-duration"
	PhenotypeNameSmellSensitivityForMalt                   PhenotypeName = "smell-sensitivity-for-malt"
	PhenotypeNameSmokingBehavior                           PhenotypeName = "smoking-behavior"
	PhenotypeNameVisceralAndSubcutaneousAdiposeTissueRatio PhenotypeName = "visceral-and-subcutaneous-adipose-tissue-ratio"
	PhenotypeNameVitaminA                                  PhenotypeName = "vitamin-a"
	PhenotypeNameVitaminB12                                PhenotypeName = "vitamin-b12"
	PhenotypeNameVitaminD                                  PhenotypeName = "vitamin-d"
	PhenotypeNameVitaminE                                  PhenotypeName = "vitamin-e"
	PhenotypeNameWaist                                     PhenotypeName = "waist"
	PhenotypeNameWaistHipRatio                             PhenotypeName = "waist-hip-ratio"
	PhenotypeNameGeneticWeight                             PhenotypeName = "weight"
	PhenotypeNameWhiteWineLiking                           PhenotypeName = "white-wine-liking"
	PhenotypeNameWordReadingAbility                        PhenotypeName = "word-reading-ability"

	// Disease

	PhenotypeNameLungCancer           PhenotypeName = "lung-cancer"
	PhenotypeNameColorectalCancer     PhenotypeName = "colorectal-cancer"
	PhenotypeNameGastricCancer        PhenotypeName = "gastric-cancer"
	PhenotypeNameBreastCancer         PhenotypeName = "breast-cancer"
	PhenotypeNameLiverCancer          PhenotypeName = "liver-cancer"
	PhenotypeNamePancreaticCancer     PhenotypeName = "pancreatic-cancer"
	PhenotypeNameProstateCancer       PhenotypeName = "prostate-cancer"
	PhenotypeNameType2Diabetes        PhenotypeName = "type-2-diabetes"
	PhenotypeNameMyocardialInfarction PhenotypeName = "myocardial-infarction"
	PhenotypeNameNicotineDependence   PhenotypeName = "nicotine-dependence"
)
