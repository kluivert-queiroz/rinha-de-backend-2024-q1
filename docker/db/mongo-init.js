let error = true;

db.customer.drop();
db.customer.insertMany([
  { _id: "1", balance: 0, limit: 100000, transactions: [], version: 1 },
  { _id: "2", balance: 0, limit: 80000, transactions: [], version: 1 },
  { _id: "3", balance: 0, limit: 1000000, transactions: [], version: 1 },
  { _id: "4", balance: 0, limit: 10000000, transactions: [], version: 1 },
  { _id: "5", balance: 0, limit: 500000, transactions: [], version: 1 },
]);
db.customer.createIndex({
  _id: 1,
  version: 1,
});
