// MongoDB Migration Script
// Este script cria os índices necessários para a coleção de orders

db = db.getSiblingDB('orders_db');

// Cria a coleção se não existir
db.createCollection('orders');

// Cria índices
db.orders.createIndex({ "customer_id": 1 });
db.orders.createIndex({ "created_at": -1 });
db.orders.createIndex({ "customer_id": 1, "created_at": -1 });

print("✅ Índices criados com sucesso!");
