local function main()
	local nodes = {}
	local node_pattern = "(%w%w%w)"
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

	local start_nodes = {}

	for node, _ in pairs(nodes) do
		if string.match(node, "%w%wA") ~= nil then
			start_nodes[#start_nodes+1] = node
		end
	end

	local n_steps = {}

	for _, node in pairs(start_nodes) do
		local n_moves = 0
		local curr_move_ind = 1
		while string.match(node, "(%u%uZ)") == nil do
			local curr_move = moves:sub(curr_move_ind, curr_move_ind)
			node = nodes[node][curr_move]
			n_moves = n_moves + 1
			curr_move_ind = curr_move_ind + 1
			if curr_move_ind > #moves then curr_move_ind = curr_move_ind - #moves end
		end
		n_steps[#n_steps+1] = n_moves
	end

	print(table.concat(n_steps, ", "))

	print(require("./scm").scm_var(n_steps))
end

main()
