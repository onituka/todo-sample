USE sample_db;

-- todos table test data
INSERT INTO todos
  (id, title, memo, implementation_date, due_date, priorities_id, complete_flag)
VALUES
  (1, "国民年金申請", "年金手帳持っていく", "2021-06-17", "2021-06-17", 3, true),
  (2, "コンセントを購入", "500円", "2021-06-17", "2021-06-17", 1, true),
  (3, "机作る", null, "2021-06-01", "2021-06-17", 1, true),
  (4, "国保を納める", null, "2021-06-25", "2021-06-25", 2, false);
