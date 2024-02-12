let error = true

db.wallet.drop()
db.wallet.insert({_id: "1", balance: 0, limit: 100000, transactions: []})
db.wallet.insert({_id: "2", balance: 0, limit: 80000, transactions: []})
db.wallet.insert({_id: "3", balance: 0, limit: 1000000, transactions: []})
db.wallet.insert({_id: "4", balance: 0, limit: 10000000, transactions: []})
db.wallet.insert({_id: "5", balance: 0, limit: 500000, transactions: []})
