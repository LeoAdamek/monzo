package monzo

import (
    "testing"
    "sort"
)

var transactions = []Transaction{
    {
        ID: "1",
        Amount: 100,
        Description: "Test A",
        Merchant: Merchant{
            Name: "M0",
        },
    },
    
    {
        ID: "2",
        Amount: 50,
        Description: "Test B",
        Merchant: Merchant{
            Name: "M1",
        },
    },
    
    {
        ID: "3",
        Amount: 200,
        Description: "Test C",
        Merchant: Merchant{
            Name: "M3",
        },
    },
}

func TestByValue(t *testing.T) {
    expected := []string{"2","1","3"}
    
    var tx []Transaction
    
    copy(transactions, tx)
    
    sort.Sort(ByValue(tx))
    
    for idx := range tx {
        if tx[idx].ID != expected[idx] {
            t.Errorf("Expected (ID = %d) as pos %d, got (ID = %d", expected[idx], idx, tx[idx].ID)
        }
    }
}