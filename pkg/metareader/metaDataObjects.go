package metareader

type Document struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"Document"`
}

type BusinessProcess struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"BusinessProcess"`
}

type Catalog struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"Catalog"`
}

type Report struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"Report"`
}

type DataProcessor struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"DataProcessor"`
}
type InformationRegister struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"InformationRegister"`
}

type AccumulationRegister struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"AccumulationRegister"`
}

type AccountingRegister struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"AccountingRegister"`
}

type ChartOfCalculationTypes struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"ChartOfCalculationTypes"`
}

type Task struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"Task"`
}

type Cube struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"Cube"`
}

type ExternalDataSource struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Table []string `xml:"Table"`
			Cube  []string `xml:"Cube"`
		}
	} `xml:"ExternalDataSource"`
}

type CalculationRegister struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form          []string `xml:"Form"`
			Recalculation []string `xml:"Recalculation"`
			Command       []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"CalculationRegister"`
}

type ChartOfCharacteristicTypes struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"ChartOfCharacteristicTypes"`
}

type ChartOfAccounts struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"ChartOfAccounts"`
}

type ExchangePlan struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"ExchangePlan"`
}

type DocumentJournal struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"DocumentJournal"`
}

type Enum struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"Enum"`
}

type CommonModuleObj struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"CommonModule"`
}

type SettingsStorage struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form []string `xml:"Form"`
		}
	} `xml:"SettingsStorage"`
}

type Table struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"Table"`
}

type FilterCriterion struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form    []string `xml:"Form"`
			Command []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"FilterCriterion"`
}

type WebService struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"WebService"`
}

type HTTPService struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"HTTPService"`
}

type Form struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"Form"`
}

type CommonForm struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"CommonForm"`
}

type CommonCommand struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"CommonCommand"`
}

type Constant struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"Constant"`
}

const (
	CommandModule             = "078a6af8-d22c-4248-9c33-7e90075a3d2c"
	ObjectModule              = "a637f77f-3840-441d-a1c3-699c8c5cb7e0"
	ManagerModule             = "d1b64a2c-8078-4982-8190-8f81aefda192"
	ManagerModuleStorage      = "0c8cad23-bf8c-468e-b49e-12f1927c048b"
	FormModule                = "32e087ab-1491-49b6-aba7-43571b41ac2b"
	RecordSetModule           = "9f36fd70-4bf4-47f6-b235-935f73aab43f"
	ValueManagerModule        = "3e58c91f-9aaa-4f42-8999-4baf33907b75"
	ManagedApplicationModule  = "d22e852a-cf8a-4f77-8ccb-3548e7792bea"
	SessionModule             = "9b7bbbae-9771-46f2-9e4d-2489e0ffc702"
	ExternalConnectionModule  = "a4a9c1e2-1e54-4c7f-af06-4ca341198fac"
	OrdinaryApplicationModule = "a78d9ce3-4e0c-48d5-9863-ae7342eedf94"
	SimpleModule              = "d5963243-262e-4398-b4d7-fb16d06484f6"
)

type coverData struct {
	Data map[Module]string
}

type Module struct {
	ModuleUuid string
	ModuleType string
}
