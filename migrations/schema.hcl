table "articles" {
  schema = schema.monolith
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "title" {
    null = true
    type = varchar(255)
  }
  column "author_id" {
    null     = true
    type     = int
    unsigned = true
  }
  primary_key {
    columns = [column.id]
  }
}
schema "monolith" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
