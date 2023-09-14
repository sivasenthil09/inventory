package controllers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ExtractCustomerID(jwtToken string, secretKey string) (int64, error) {

	// Parse the JWT token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err // Handle token parsing errors
	}

	// Check if the token is valid
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Extract the customer ID from the claims
			customerID, ok := claims["customerid"].(int64)
			if ok {
				return customerID, nil
			}
		}
	}

	return 0, fmt.Errorf("Invalid or expired JWT token")
}
