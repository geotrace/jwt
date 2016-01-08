// jwt поддерживает работу с токенами в формате JSON Web Token.
//
// В данной библиотеке я решил ограничиться только поддержкой основного метода подписи — HS256.
// Во-первых, потому что он единственный является обязательным. Во-вторых, совершенно не хотелось
// подключать дополнительные библиотеки без необходимости. В-третьих, с ним получается одна из
// самых коротких подписей, что, как вы понимаете, в прямую касается длины получаемого токена.
// В общем, мне показалось этого достаточно.
package jwt
