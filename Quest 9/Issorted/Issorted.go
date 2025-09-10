package student

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 { // slice vide ou 1 élément → toujours trié
		return true
	}

	direction := 0 // 1 = croissant, -1 = décroissant, 0 = pas encore défini

	for i := 1; i < len(a); i++ { // parcours du slice
		cmp := f(a[i-1], a[i]) // compare l'élément précédent et le courant
		if cmp != 0 {          // on ignore les égalités
			if direction == 0 { // on définit le sens du tri
				if cmp > 0 {
					direction = -1 // décroissant
				} else {
					direction = 1 // croissant
				}
			} else { // sens déjà défini, on vérifie cohérence
				if (cmp > 0 && direction == 1) || (cmp < 0 && direction == -1) {
					return false // violation du tri
				}
			}
		}
	}

	return true // aucune violation trouvée → trié
}
