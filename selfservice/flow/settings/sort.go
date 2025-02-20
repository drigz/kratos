package settings

import (
	"context"

	"github.com/ory/kratos/ui/node"
)

var ctx = context.Background()

func sortNodes(n node.Nodes, schemaRef string) error {
	return n.SortBySchema(ctx,
		node.SortBySchema(schemaRef),
		node.SortByGroups([]node.Group{
			node.DefaultGroup,
			node.ProfileGroup,
			node.PasswordGroup,
			node.OpenIDConnectGroup,
			node.LookupGroup,
			node.WebAuthnGroup,
			node.TOTPGroup,
		}),
		node.SortUseOrderAppend([]string{
			// Lookup
			node.LookupReveal,
			node.LookupRegenerate,
			node.LookupCodes,
			node.LookupConfirm,

			// Lookup
			node.WebAuthnRemove,
			node.WebAuthnRegisterDisplayName,
			node.WebAuthnRegister,

			// TOTP
			node.TOTPQR,
			node.TOTPSecretKey,
			node.TOTPUnlink,
			node.TOTPCode,
		}),
	)
}
