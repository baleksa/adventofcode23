local function gcd(x, y)
	if y == 0 then
		return x
	else
		return gcd(y, x % y)
	end
end

local function scm(x, y)
	return x * y // gcd(x, y)
end

local function scm_var(nums)
	local result = 1
	for _, v in ipairs(nums) do
		result = scm(v, result)
	end
	return result
end

return {
	gcd = gcd,
	scm = scm,
	scm_var = scm_var,
}
