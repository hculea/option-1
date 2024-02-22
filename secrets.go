// AUTOGENERATED - DO NOT EDIT MANUALLY
package onepassword

type SecretsAPI interface {
	 Resolve(secretReference string) (string, error) 
	
}

type SecretsSource struct {
	InnerClient
}

func NewSecretsSource(inner InnerClient) *SecretsSource {
	return &SecretsSource{inner}
}


func (s SecretsSource) Resolve(secretReference string) (string, error) { 
	res, err := clientInvoke(s.InnerClient, "Resolve", map[string]interface{}{
		"secret_reference": secretReference,
	})
	if err != nil {
		return "", err
	}
	return *res, nil
}
