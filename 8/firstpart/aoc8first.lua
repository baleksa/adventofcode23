local L <const> = "L"
local R <const> = "R"
local START <const> = "AAA"
local END <const> = "ZZZ"
local nodes = {}
local node_pattern = "(%u%u%u)"
local line_pattern = string.format("%s = %%(%s, %s%%)", node_pattern, node_pattern, node_pattern)
local moves
for line in io.lines(arg[1]) do
	if moves == nil then
		moves = string.match(line, "([LR]+)")
		goto continue
	end
	if #line < 9 then
		goto continue
	end
	local node, left, right = string.match(line, line_pattern)

	nodes[node] = {L=left, R=right}
	::continue::
end

-- for node, conn in pairs(nodes) do
-- 	print(node, conn[L], conn[R])
-- end

local curr_node = START
local n_moves = 0
local curr_move_ind = 1
while curr_node ~= END do
	local curr_move = moves:sub(curr_move_ind, curr_move_ind)
	curr_node = nodes[curr_node][curr_move]
	n_moves = n_moves + 1
	curr_move_ind = curr_move_ind + 1
	if curr_move_ind > #moves then curr_move_ind = curr_move_ind - #moves end
end

print(n_moves)
