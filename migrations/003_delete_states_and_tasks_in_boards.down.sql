ALTER TABLE states DROP CONSTRAINT IF EXISTS states_board_id_fkey;
ALTER TABLE states ADD CONSTRAINT states_board_id_fkey FOREIGN KEY (board_id) REFERENCES boards(id);

ALTER TABLE tasks DROP CONSTRAINT IF EXISTS tasks_state_id_fkey;
ALTER TABLE tasks ADD CONSTRAINT tasks_state_id_fkey FOREIGN KEY (state_id) REFERENCES states(id);
