package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// GET /api/projects/all → []Project
type Project struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Year        string   `json:"year"` // e.g. "2024" — optional
	URL         string   `json:"url"`  // optional
}

var AllProjects = []Project{
  {
    Title:       "AffineAlign-Py",
    Description: "A high-performance bioinformatics tool implementing the Affine Gap Penalty algorithm for precise global and local sequence alignment, optimized for biological accuracy.",
    Tags:        []string{"Bioinformatics", "Algorithms", "Python", "BBL434"},
    Year:        "2026",
    URL:         "https://github.com/Dhruv-78/BBL434-Assignment-Affine-Gapped-sequence-alignment-toolAssignment.git",
  },
  {
    Title:       "PlasmidForge",
    Description: "An automated sequence assembly utility designed for the construction and mapping of universal plasmids in molecular biology workflows.",
    Tags:        []string{"Biotechnology", "Automation", "Molecular Biology", "Genetic Engineering"},
    Year:        "2026",
    URL:         "https://github.com/Dhruv-78/Assignment-Universal-Plasmid-Maker.git",
  },
  {
    Title:       "FindMyFace",
    Description: "A computer vision application utilizing spatial detection frameworks to identify and localize human faces within complex image backgrounds.",
    Tags:        []string{"Computer Vision", "OpenCV", "Object Detection", "Python"},
    Year:        "2024",
    URL:         "https://github.com/Dhruv-78/Face-finder.git",
  },
  {
    Title:       "RetailPulse",
    Description: "A consumer behavior engine that analyzes historical transaction data to predict future purchase propensity and market trends.",
    Tags:        []string{"E-commerce", "Behavioral Analytics", "Machine Learning", "Python"},
    Year:        "2023",
    URL:         "https://github.com/Dhruv-78/Purchase_prediction.git",
  },
  {
    Title:       "Sentinelle-SMS",
    Description: "An NLP-driven classification system designed to filter and protect against mobile communication spam using text vectorization.",
    Tags:        []string{"NLP", "Classification", "Security", "Scikit-Learn"},
    Year:        "2023",
    URL:         "https://github.com/Dhruv-78/SMS-Spam-Ham-Prediction.git",
  },
  {
    Title:       "CardioScan",
    Description: "A healthcare diagnostic model that evaluates patient clinical metrics to provide high-accuracy risk assessments for heart disease.",
    Tags:        []string{"Healthcare AI", "Predictive Modeling", "Data Science", "Medical Tech"},
    Year:        "2025",
    URL:         "https://github.com/Dhruv-78/Heart_disease_prediction.git",
  },
  {
    Title:       "VitalMetrics-Analyzer",
    Description: "A global health regression model that correlates socioeconomic indicators with average life expectancy to project demographic shifts.",
    Tags:        []string{"Regression", "Data Analytics", "Public Health", "Statistics"},
    Year:        "2025",
    URL:         "https://github.com/Dhruv-78/life_expectancy_prediction.git",
  },
  {
    Title:       "InsuRate",
    Description: "An insurance tech regression engine developed to estimate individual medical costs based on historical risk factors and demographics.",
    Tags:        []string{"InsurTech", "Regression", "Machine Learning", "Fintech"},
    Year:        "2025",
    URL:         "https://github.com/Dhruv-78/Medical_cost_prediction.git",
  },
  {
    Title:       "PropVal-Estimator",
    Description: "A real estate valuation model that utilizes multi-feature regression to predict housing market prices with high precision.",
    Tags:        []string{"Real Estate AI", "Regression", "Predictive Analytics", "Data Science"},
    Year:        "2023",
    URL:         "https://github.com/Dhruv-78/House_price.git",
  },
  {
    Title:       "PokeClassify",
    Description: "A multi-class classification project exploring feature engineering to predict elemental types based on character statistics.",
    Tags:        []string{"Classification", "Feature Engineering", "Python", "Data Science"},
    Year:        "2024",
    URL:         "https://github.com/Dhruv-78/Pokemon-type-guesser.git",
  },
  {
    Title:       "Tasksify",
    Description: "A productivity-focused task management application designed with a robust backend to handle full CRUD operations for workflow optimization.",
    Tags:        []string{"Software Development", "Backend", "CRUD", "Productivity"},
    Year:        "2026",
    URL:         "https://github.com/Dhruv-78/Tasksify.git",
  },
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	projects := AllProjects
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Year > projects[j].Year
	})
	projects = projects[:4]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Logic to send an email or save to a database goes here
	fmt.Fprint(w, "Message received successfully!")
}

func getAllProjects(w http.ResponseWriter, r *http.Request) {
	// This would typically fetch from a database, but we'll return a static list for now
	allProjects := AllProjects
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allProjects)
}
